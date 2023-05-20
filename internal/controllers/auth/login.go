package auth

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/middlewares"
)

// Login
// @Tags User
// @Description 登录
// @Accept json
// @Product json
// @Param login body models.LoginModel true "Login Model"
// @Router /api/login [post]
// @Success 200
func Login(ctx *gin.Context) {
	middlewares.GinJwtMiddleware().LoginHandler(ctx)
}
