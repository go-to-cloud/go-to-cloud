package auth

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/users"
	"net/http"
)

// GetAuthCodes
// @Tags User
// @Description 获取当前用户拥有的权限点
// @Accept json
// @Product json
// @Router /api/user/auths [get]
// @Success 200
func GetAuthCodes(ctx *gin.Context) {
	exists, _, _, _, _, kinds := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	codes := users.GetAuthCodes(kinds)

	response.Success(ctx, codes)
}
