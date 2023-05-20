package middlewares

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	casbinAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models"
	repo "go-to-cloud/internal/repositories"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

type login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var jwtMiddleware *jwt.GinJWTMiddleware

func GinJwtMiddleware() *jwt.GinJWTMiddleware {
	return jwtMiddleware
}

func skip(ctx *gin.Context) bool {
	return ctx.Request.Method == http.MethodPut && strings.EqualFold(ctx.FullPath(), "/api/projects/:projectId/deploy/:id") // 重新部署允许guest操作
}

func AuthHandler(enforcer *casbin.Enforcer) gin.HandlerFunc {
	m, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       conf.GetJwtKey().Realm,
		Key:         []byte(conf.GetJwtKey().Security),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour / 2,
		IdentityKey: "jti",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*repo.User); ok {
				orgs := make(map[int64]string, len(v.Orgs))
				for _, org := range v.Orgs {
					orgs[int64(org.ID)] = org.Name
				}
				return jwt.MapClaims{
					"jti":  v.ID,
					"sub":  v.Account,
					"kind": v.Kind,
					"orgs": orgs,
				}
			}
			return jwt.MapClaims{}
		},
		// 认证
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVal login
			if err := c.ShouldBind(&loginVal); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			account := loginVal.UserName
			password := loginVal.Password

			user := repo.GetUser(&account, &password)

			if user != nil {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		// 鉴权
		Authorizator: func(data interface{}, c *gin.Context) bool {
			claims, _ := jwtMiddleware.GetClaimsFromJWT(c)
			ok, ex := enforcer.Enforce(claims["sub"], c.Request.URL.Path, c.Request.Method)
			if ok {
				return true
			}

			if claims["kind"] == nil {
				return false
			}
			kind := claims["kind"].([]interface{})
			for _, s := range kind {

				if models.Kind(s.(string)) == models.Guest && (c.Request.Method == http.MethodPut || c.Request.Method == http.MethodDelete) {
					if !skip(c) {
						return false
					}
				}

				ok, ex = enforcer.Enforce(s, c.Request.URL.Path, c.Request.Method)
				if ex != nil {
					log.Fatal(ex)
				}
				if ok {
					return true
				}
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		HTTPStatusMessageFunc: func(e error, c *gin.Context) string {
			return "请见谅，游客不允许此操作，如为误添加，请联系993921@qq.com协助"
		},
	})

	jwtMiddleware = m

	return m.MiddlewareFunc()
}

func GetCasbinEnforcer(db *gorm.DB) (*casbin.Enforcer, error) {
	adapter, err := casbinAdapter.NewAdapterByDBWithCustomTable(db, nil, "casbin_rules")
	if err != nil {
		return nil, err
	}

	rbacModel, err := model.NewModelFromString(fmt.Sprintf(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = (g(r.sub, p.sub) && (keyMatch4(r.obj,p.obj) || keyMatch5(r.obj,p.obj)) && (r.act == p.act)) || r.sub == "%s"
`, models.RootUserName))

	return casbin.NewEnforcer(rbacModel, adapter)
}
