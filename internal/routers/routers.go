package routers

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/conf"
	auth2 "go-to-cloud/internal/auth"
	"go-to-cloud/internal/controllers/auth"
	"go-to-cloud/internal/controllers/users"
	"go-to-cloud/internal/middlewares"
	"strings"
)

// buildRouters 构建路由表
func buildRouters(router *gin.Engine) {

	api := router.Group("/api")
	api.POST("/login", auth.Login)
	api.GET("/user/logout", users.Logout)

	enforcer, err := middlewares.GetCasbinEnforcer(conf.GetDbClient())
	if err != nil {
		panic(err)
	}
	router.Use(middlewares.AuthHandler(enforcer))
	{
		for _, routerMap := range auth2.RouterMaps {
			if strings.HasPrefix(routerMap.Url, "/api") {
				for _, method := range routerMap.Methods {
					if method == auth2.GET {
						router.GET(routerMap.Url, routerMap.Func)
					} else if method == auth2.PUT {
						router.PUT(routerMap.Url, routerMap.Func)
					} else if method == auth2.POST {
						router.POST(routerMap.Url, routerMap.Func)
					} else if method == auth2.DELETE {
						router.DELETE(routerMap.Url, routerMap.Func)
					} // ignore rest methods
				}
			}
		}
	}
}

func buildWebSocket(router *gin.Engine) {
	for _, routerMap := range auth2.RouterMaps {
		if strings.HasPrefix(routerMap.Url, "/ws") {
			for _, method := range routerMap.Methods {
				if method == auth2.GET {
					router.GET(routerMap.Url, routerMap.Func)
				}
			}
		}
	}
}
