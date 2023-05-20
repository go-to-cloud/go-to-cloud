package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/utils"
	"os"
	"testing"
)

func TestMany2ManyMemberToOrg(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		if conf.GetDbClient().Migrator().HasTable(Org{}) {
			conf.GetDbClient().Migrator().DropTable(&Org{})
		}
		conf.GetDbClient().Migrator().AutoMigrate(&Org{})
		if conf.GetDbClient().Migrator().HasTable(User{}) {
			conf.GetDbClient().Migrator().DropTable(&User{})
		}
		conf.GetDbClient().Migrator().AutoMigrate(&User{})
	}

	u1 := user.User{
		Account:        "u1",
		RealName:       realName,
		OriginPassword: pwd,
		Email:          email,
		Mobile:         mobile,
	}
	u2 := user.User{
		Account:        "u2",
		RealName:       realName,
		OriginPassword: pwd,
		Email:          email,
		Mobile:         mobile,
	}
	err := CreateUser(&u1)
	assert.NoError(t, err)
	err = CreateUser(&u2)
	assert.NoError(t, err)

	err = CreateOrg(&orgName, &orgRemark)
	assert.NoError(t, err)
	org, err := GetOrgs()
	assert.NoError(t, err)

	err = UpdateMembersToOrg(org[0].ID, []uint{u1.Id, u2.Id}, []uint{})
	assert.NoError(t, err)

	afterOrg, err := GetOrgs()
	assert.NoError(t, err)
	assert.Len(t, afterOrg[0].Users, 2)

	err = UpdateMembersToOrg(org[0].ID, []uint{}, []uint{u1.Id})
	users, err := GetUsersByOrg(org[0].ID)
	assert.NoError(t, err)
	o := utils.New[uint]()
	for i := range users {
		utils.Add(o, users[i].ID)
	}
	assert.True(t, utils.Has(o, u2.Id))
	assert.False(t, utils.Has(o, u1.Id))
}

func TestMany2ManyOrgToUser(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		if conf.GetDbClient().Migrator().HasTable(Org{}) {
			conf.GetDbClient().Migrator().DropTable(&Org{})
		}
		conf.GetDbClient().Migrator().AutoMigrate(&Org{})
		if conf.GetDbClient().Migrator().HasTable(User{}) {
			conf.GetDbClient().Migrator().DropTable(&User{})
		}
		conf.GetDbClient().Migrator().AutoMigrate(&User{})
	}

	u := user.User{
		Account:        "u2",
		RealName:       realName,
		OriginPassword: pwd,
		Email:          email,
		Mobile:         mobile,
	}
	err := CreateUser(&u)
	assert.NoError(t, err)

	err = CreateOrg(&orgName, &orgRemark)
	assert.NoError(t, err)
	err = CreateOrg(&orgName, &orgRemark)
	assert.NoError(t, err)

	org, err := GetOrgs()
	assert.NoError(t, err)

	err = UpdateOrgsToMember(u.Id, []uint{org[0].ID, org[1].ID}, []uint{})
	assert.NoError(t, err)

	afterOrgs, err := GetOrgsByUser(u.Id)
	assert.NoError(t, err)

	o := utils.New[uint]()
	for _, ao := range afterOrgs {
		utils.Add(o, ao.ID)
	}
	assert.True(t, utils.Has(o, org[0].ID))
	assert.True(t, utils.Has(o, org[1].ID))

	err = UpdateOrgsToMember(u.Id, []uint{}, []uint{org[1].ID})
	assert.NoError(t, err)

	afterOrgs, err = GetOrgsByUser(u.Id)
	assert.NoError(t, err)

	o = utils.New[uint]()
	for _, ao := range afterOrgs {
		utils.Add(o, ao.ID)
	}
	assert.True(t, utils.Has(o, org[0].ID))
	assert.False(t, utils.Has(o, org[1].ID))
}
