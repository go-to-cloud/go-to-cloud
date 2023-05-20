package repositories

import (
	"errors"
	"go-to-cloud/conf"
	"gorm.io/gorm/clause"
)

// Org 组织
// 组织拥有基础设施和用户
type Org struct {
	Model
	Name   string  `json:"name" gorm:"column:name;not null;type:nvarchar(20)"` // 组织名称
	Users  []*User `gorm:"many2many:orgs_users_rel;"`
	Remark string  `json:"remark" gorm:"column:remark;type:nvarchar(1024);"`
}

func (m *Org) TableName() string {
	return "org"
}

func GetOrgs() ([]Org, error) {
	db := conf.GetDbClient()

	var org []Org
	err := db.Preload("Users").Find(&org).Error

	return org, err
}

func CreateOrg(name, remark *string) error {
	if name == nil || len(*name) == 0 {
		return errors.New("组织名称不能为空")
	}
	db := conf.GetDbClient()

	org := &Org{
		Name:   *name,
		Remark: *remark,
	}
	err := db.Model(&Org{}).Create(org).Error

	return err
}

func UpdateOrg(id uint, name, remark *string) error {
	if len(*name) == 0 {
		return errors.New("组织名称不能为空")
	}

	db := conf.GetDbClient()

	org := &Org{
		Name:   *name,
		Remark: *remark,
	}
	err := db.Model(&Org{}).Where("id = ?", id).Select("name", "remark").Updates(org).Error

	return err
}

func DeleteOrg(id uint) error {
	db := conf.GetDbClient()

	var org Org
	err := db.Model(&Org{}).Preload(clause.Associations).First(&org, id).Error
	if err != nil {
		return err
	}

	if len(org.Users) > 0 {
		return errors.New("组织中存在用户，请先移除所有用户后再删除组织")
	}

	var total int64
	err = db.Model(&Org{}).Where("id != ?", id).Count(&total).Error
	if err != nil {
		return err
	}
	if total == 0 {
		return errors.New("至少需要保留一个组织")
	}

	err = db.Delete(&org).Error
	return err
}
