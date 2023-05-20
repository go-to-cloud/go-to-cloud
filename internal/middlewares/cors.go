package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsHandler 跨域处理中间件
func CorsHandler() gin.HandlerFunc {
	return cors.Default()
}
