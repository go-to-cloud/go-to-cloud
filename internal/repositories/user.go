package repositories

import (
	"encoding/json"
	"errors"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
)

// User 登录账户
type User struct {
	Model
	RealName       string         `json:"realName" gorm:"column:real_name;type:nvarchar(16);not null;default('')"`
	Pinyin         string         `json:"pinyin" gorm:"column:pinyin;type:nvarchar(100);not null;default('')"`           // RealName的拼音
	PinyinInit     string         `json:"pinyin_init" gorm:"column:pinyin_init;type:nvarchar(100);not null;default('')"` // RealName的拼音首字母
	Account        string         `json:"account" gorm:"column:account;not null;type:nvarchar(200)"`                     // 账号
	HashedPassword string         `json:"-" gorm:"column:password;not null;type:nvarchar(200)"`                          // 登录密码
	Email          string         `json:"email" gorm:"column:email;type:nvarchar(200)"`                                  // 邮箱
	Mobile         string         `json:"mobile" gorm:"column:mobile;type:nvarchar(200)"`                                // 联系电话
	LastLoginAt    *time.Time     `json:"last_login_at" gorm:"column:last_login_at"`                                     // 上次登录时间
	Kind           datatypes.JSON `json:"kind" gorm:"column:kind;"`
	Orgs           []*Org         `gorm:"many2many:orgs_users_rel"`
}

func (m *User) TableName() string {
	return "users"
}
func (m *User) BeforeDelete(_ *gorm.DB) (err error) {
	if strings.EqualFold("root", m.Account) {
		return errors.New("root用户无法被删除")
	}
	return
}

// SetPassword 加密密码
func (m *User) SetPassword(origPassword *string) error {
	if len(strings.Trim(*origPassword, " ")) == 0 {
		return errors.New("密码不允许为空")
	}
	lowerPassword := strings.ToLower(*origPassword)
	if hashBytes, err := bcrypt.GenerateFromPassword([]byte(lowerPassword), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		m.HashedPassword = string(hashBytes)
		return nil
	}
}

// GetUser by account AND password
func GetUser(account, password *string) *User {
	tx := conf.GetDbClient()

	var u User

	if tx.Preload(clause.Associations).Where(&User{Account: *account}).First(&u).Error != nil {
		return nil
	}
	if u.comparePassword(password) {
		return &u
	}
	return nil
}

// comparePassword 比较密码
func (m *User) comparePassword(password *string) bool {
	lowerPassword := strings.ToLower(*password)
	return nil == bcrypt.CompareHashAndPassword([]byte(m.HashedPassword), []byte(lowerPassword))
}

// GetAllUser 获取所有用户
func GetAllUser() ([]User, error) {
	tx := conf.GetDbClient()

	var users []User

	err := tx.Preload(clause.Associations).Find(&users).Error

	return users, err
}

func valid(id uint, user *user.User) error {
	if len(user.Account) == 0 {
		return errors.New("账号名称不能为空")
	}

	if len(user.Email) > 0 {
		if !utils.IsValidEmail(user.Email) {
			return errors.New("邮箱格式不正确")
		}
	}

	if len(user.Mobile) > 0 {
		if !utils.IsValidMobile(user.Mobile) {
			return errors.New("手机号码格式不正确")
		}
	}

	db := conf.GetDbClient()
	var tar User
	tx := db.Model(&User{}).Where("account = ?", user.Account)
	if id > 0 {
		tx.Where("id != ?", id)
	}
	err := tx.First(&tar).Error

	if err == nil {
		return errors.New("账号已经存在")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

func mapper(user *user.User, ignorePwd bool) (*User, error) {
	user.TransPinyin()
	u := &User{
		RealName:   user.RealName,
		Pinyin:     user.Pinyin,
		PinyinInit: user.PinyinInit,
		Account:    user.Account,
		Email:      user.Email,
		Mobile:     user.Mobile,
		Kind: func() datatypes.JSON {
			k, _ := json.Marshal(user.Kind)
			return k
		}(),
	}
	var err error
	if !ignorePwd {
		err = u.SetPassword(&user.OriginPassword)
	}
	return u, err
}

func CreateUser(user *user.User) error {
	if err := valid(0, user); err != nil {
		return err
	}

	db := conf.GetDbClient()

	if u, err := mapper(user, false); err != nil {
		return err
	} else {
		if err := db.Model(&User{}).Create(u).Error; err == nil {
			user.Id = u.ID
			return nil
		} else {
			return err
		}
	}
}

func UpdateUser(id uint, user *user.User) error {
	if id == 0 {
		return errors.New("找不到用户")
	}

	if err := valid(id, user); err != nil {
		return err
	}

	db := conf.GetDbClient()

	if u, err := mapper(user, true); err != nil {
		return err
	} else {
		return db.Model(&User{}).Where("id = ?", id).Omit("account", "password").Updates(u).Error
	}
}

func DeleteUser(id uint) error {
	db := conf.GetDbClient()

	var total int64
	err := db.Model(&User{}).Where("id != ?", id).Count(&total).Error
	if err != nil {
		return err
	}
	if total == 0 {
		return errors.New("至少需要保留一个用户")
	}

	var targetUser User
	db.Model(&User{}).Find(&targetUser, id)
	err = db.Delete(&targetUser, id).Error
	return err
}

func ResetPassword(id uint, origPassword *string) error {
	db := conf.GetDbClient()

	u := &User{Model: Model{ID: id}}
	u.SetPassword(origPassword)
	err := db.Select("password").Updates(u).Error
	if err != nil {
		return err
	}
	return err
}

func ResetPasswordWithCheckOldPassword(id uint, oldPassword, origPassword *string) error {
	db := conf.GetDbClient()

	var u User
	if err := db.First(&u, id).Error; err != nil {
		return err
	}

	if !u.comparePassword(oldPassword) {
		return errors.New("旧密码不正确")
	}

	u.SetPassword(origPassword)
	err := db.Select("password").Updates(u).Error
	if err != nil {
		return err
	}
	return err
}
