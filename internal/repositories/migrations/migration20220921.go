package migrations

import (
	"github.com/casbin/casbin/v2"
	"go-to-cloud/internal/auth"
	"go-to-cloud/internal/middlewares"
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
	"regexp"
)

type migration20220921 struct {
}

// addGroupPolicy 添加角色继承关系
func addGroupPolicy(enforce *casbin.Enforcer) {
	for _, strings := range auth.GroupPolicies() {
		if _, err := enforce.AddGroupingPolicy(strings[0], strings[1]); err != nil {
			panic(err)
		}
	}
}

// addResourcePolicy 添加权限点
func addResourcePolicy(enforce *casbin.Enforcer) {
	for _, p := range auth.ResourcePolicies() {
		if _, err := enforce.AddPolicies(p); err != nil {
			panic(err)
		}
	}
}

// addRouterPolicy 添加路由权限
func addRouterPolicy(enforce *casbin.Enforcer) {
	reg := regexp.MustCompile(`:(\w+)`)
	for _, routerMap := range auth.RouterMaps {
		for _, kind := range routerMap.Kinds {
			for _, method := range routerMap.Methods {
				// 需要将路由参数 :params 替换为 {params}来适配keyMatch4匹配算法
				if _, err := enforce.AddPolicies([][]string{{string(kind), reg.ReplaceAllString(routerMap.Url, "{$1}"), string(method)}}); err != nil {
					panic(err)
				}
			}
		}
	}
}

func (m *migration20220921) Up(db *gorm.DB) error {

	if !db.Migrator().HasTable(&repo.CasbinRule{}) {
		err := db.AutoMigrate(&repo.CasbinRule{})
		if err != nil {
			return err
		} else {
			if enforce, err := middlewares.GetCasbinEnforcer(db); err == nil {
				addGroupPolicy(enforce)
				addResourcePolicy(enforce)
				addRouterPolicy(enforce)
			} else {
				return err
			}
		}
	}
	return nil
}

func (m *migration20220921) Down(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&repo.CasbinRule{},
	)
	return err
}
