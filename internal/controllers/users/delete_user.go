package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/repositories"
	"net/http"
	"strconv"
)

// DeleteUser
// @Tags User
// @Description 删除用户
// @Success 200
// @Router /api/user/{userId} [delete]
// @Security JWT
func DeleteUser(ctx *gin.Context) {
	exists, currentUserId, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var err error
	userIdStr := ctx.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	if currentUserId == uint(userId) {
		msg := "无法删除自己"
		response.Fail(ctx, http.StatusBadRequest, &msg)
		return
	}

	err = repositories.DeleteUser(uint(userId))
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
