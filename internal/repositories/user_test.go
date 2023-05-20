package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/user"
	"os"
	"testing"
	"time"
)

func TestPasswordGenAndCompare(t *testing.T) {
	u := &User{}

	pwd := "OJBK"
	u.SetPassword(&pwd)
	assert.True(t, u.comparePassword(&pwd))

	pwd = "OJBK"
	u.SetPassword(&pwd)
	assert.True(t, u.comparePassword(&pwd))

	pwd1 := "Ojbk"
	assert.True(t, u.comparePassword(&pwd1))

	pwd2 := "Ojbk "
	assert.False(t, u.comparePassword(&pwd2))

	pwd3 := ""
	assert.Error(t, u.SetPassword(&pwd3))
	assert.False(t, u.comparePassword(&pwd3))
}

const (
	account  = "root"
	realName = "肉哦"
	pwd      = "123456"
	email    = "123@email.com"
	mobile   = "13857111111"
)

func TestCreateUser(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		conf.GetDbClient().Migrator().AutoMigrate(&User{})
	}

	u := user.User{
		RealName:       realName,
		OriginPassword: pwd,
		Email:          email,
		Mobile:         mobile,
	}
	err := CreateUser(&u)
	assert.Error(t, err)

	u.Account = account
	err = CreateUser(&u)
	assert.NoError(t, err)

	all, err := GetAllUser()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(all), 1)
}

func TestUpdateUser(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		if conf.GetDbClient().Migrator().HasTable(&User{}) {
			conf.GetDbClient().Migrator().DropTable(&User{})
		}
		conf.GetDbClient().Migrator().AutoMigrate(&User{})
	}

	err := UpdateUser(0, &user.User{})
	assert.Error(t, err)

	u := user.User{
		Account:        account,
		RealName:       realName,
		OriginPassword: pwd,
		Email:          email,
		Mobile:         mobile,
	}
	err = CreateUser(&u)

	expectedUser := user.User{
		Id:             1,
		Account:        "RRRRR", // ignore when update
		RealName:       "改名了",
		OriginPassword: "654321", // ignore when update
		Email:          "321@email.com",
		Mobile:         "13857111112",
	}
	err = UpdateUser(1, &expectedUser)
	assert.NoError(t, err)

	after, err := GetAllUser()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(after), 1)

	lastUser := after[len(after)-1]
	assert.Equal(t, expectedUser.Mobile, lastUser.Mobile)
	assert.Equal(t, expectedUser.Email, lastUser.Email)
	origPwd := pwd
	assert.True(t, lastUser.comparePassword(&origPwd))
	assert.Equal(t, account, lastUser.Account)
	assert.Equal(t, expectedUser.RealName, lastUser.RealName)
}

func TestDeleteUser(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		conf.GetDbClient().Migrator().AutoMigrate(&User{})
	}

	u := user.User{
		Account:        account,
		RealName:       realName,
		OriginPassword: pwd,
		Email:          email,
		Mobile:         mobile,
	}
	err := CreateUser(&u)

	all, err := GetAllUser()
	assert.NoError(t, err)
	assert.Len(t, all, 1)

	err = DeleteUser(all[0].ID)
	assert.Error(t, err) // 至少需要保留一个用户

	err = CreateUser(&user.User{
		Account:        account + "1",
		RealName:       realName + "1",
		OriginPassword: pwd + "1",
		Email:          email + ".cn",
		Mobile:         "13857115555",
	})

	all, err = GetAllUser()
	assert.NoError(t, err)
	assert.Len(t, all, 2)

	for _, u := range all {
		err = DeleteUser(u.ID)
		if u.Account == "root" {
			assert.Error(t, err) // root用户不能删除
		} else {
			assert.NoError(t, err)
		}
	}

	after, err := GetAllUser()
	assert.NoError(t, err)
	assert.Len(t, after, 1)
}

func TestChangePassword(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		conf.GetDbClient().Migrator().AutoMigrate(&User{})
	}

	account := "test-" + time.Now().Format("20060102150405")
	pwd := "123456"
	u := &user.User{
		Account:        account,
		RealName:       "肉哦",
		OriginPassword: pwd,
		Email:          "123@email.com",
		Mobile:         "13857111111",
	}
	err := CreateUser(u)
	assert.NoError(t, err)
	u3 := GetUser(&account, &pwd)

	pwd = "34567"
	err = ResetPassword(u3.ID, &pwd)
	assert.NoError(t, err)
	u2 := GetUser(&account, &pwd)
	assert.True(t, u2.ID == u3.ID)

	pwd2 := "7777"
	wrongOldPwd := "345678"
	err = ResetPasswordWithCheckOldPassword(u3.ID, &wrongOldPwd, &pwd2)
	assert.Error(t, err)
	err = ResetPasswordWithCheckOldPassword(u3.ID, &pwd, &pwd2)
	assert.NoError(t, err)
	u4 := GetUser(&account, &pwd2)
	assert.True(t, u4.ID == u3.ID)
}
