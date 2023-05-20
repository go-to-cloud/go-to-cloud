package k8s

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strconv"
)

// RemoveK8sRepo 移除k8s仓库
// @Tags Configure
// @Description k8s仓库配置
// @Success 200
// @Router /api/configure/deploy/k8s/{id} [delete]
// @Param   id     path     int     true	"K8sRepo.ID"
// @Security JWT
func RemoveK8sRepo(ctx *gin.Context) {
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

	err = k8s.RemoveRepo(userId, uint(repoId))

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
