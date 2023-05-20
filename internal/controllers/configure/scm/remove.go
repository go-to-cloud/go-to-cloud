package scm

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
	"strconv"
)

// RemoveCodeRepo 移除代码仓库
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Router /api/configure/coderepo/{id} [delete]
// @Param   id     path     int     true	"CodeRepo.ID"
// @Security JWT
func RemoveCodeRepo(ctx *gin.Context) {
	val := ctx.Param("id")

	repoId, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = scm.RemoveRepo(userId, uint(repoId))

	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = ""
	}
	response.Success(ctx, gin.H{
		"success": err == nil,
		"message": message,
	})
}
