package repositories

import (
	"github.com/stretchr/testify/assert"
	"go-to-cloud/conf"
	"os"
	"testing"
)

var (
	orgName   = "orgName"
	orgRemark = "orgRemark"
)

func TestCreateOrg(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		conf.GetDbClient().Migrator().AutoMigrate(&Org{})
	}

	name := orgName
	remark := orgRemark

	err := CreateOrg(nil, &remark)
	assert.Error(t, err, "组织名称不能为空")

	err = CreateOrg(&name, &remark)
	assert.NoError(t, err)
}

func TestDeleteOrg(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		if conf.GetDbClient().Migrator().HasTable(Org{}) {
			conf.GetDbClient().Migrator().DropTable(&Org{})
		}
		conf.GetDbClient().Migrator().AutoMigrate(&Org{})
	}

	name := orgName
	remark := orgRemark

	err := CreateOrg(&name, &remark)
	assert.NoError(t, err)

	orgs, err := GetOrgs()
	assert.NoError(t, err)
	assert.Len(t, orgs, 1)

	err = DeleteOrg(orgs[0].ID)
	assert.Error(t, err)

	afterName := "edit" + name
	afterRemark := "edit" + remark
	err = CreateOrg(&afterName, &afterRemark)
	assert.NoError(t, err)

	orgs, err = GetOrgs()
	assert.Len(t, orgs, 2)

	err = DeleteOrg(orgs[1].ID)
	assert.NoError(t, err)
}

func TestUpdateOrg(t *testing.T) {
	if err := os.Setenv("UnitTestEnv", "1"); err != nil {
		t.Skip("skipped due to CI")
	} else {
		defer os.Unsetenv("UnitTestEnv")
		if conf.GetDbClient().Migrator().HasTable(Org{}) {
			conf.GetDbClient().Migrator().DropTable(&Org{})
		}
		conf.GetDbClient().Migrator().AutoMigrate(&Org{})
	}

	name := orgName
	remark := orgRemark

	err := CreateOrg(&name, &remark)
	assert.NoError(t, err)

	orgs, err := GetOrgs()
	assert.NoError(t, err)
	assert.Len(t, orgs, 1)

	afterName := ""
	afterRemark := ""
	err = UpdateOrg(orgs[0].ID, &afterName, &afterRemark)
	assert.Error(t, err)

	afterName = "edit" + name
	afterRemark = "edit" + remark
	err = UpdateOrg(orgs[0].ID, &afterName, &afterRemark)
	assert.NoError(t, err)

	orgs, err = GetOrgs()
	assert.NoError(t, err)
	assert.Equal(t, orgs[0].Name, afterName)
	assert.Equal(t, orgs[0].Remark, afterRemark)
}
