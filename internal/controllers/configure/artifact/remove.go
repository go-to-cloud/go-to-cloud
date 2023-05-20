package artifact

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strconv"
)

// RemoveArtifactRepo 移除制品仓库
// @Tags Configure
// @Description 制品仓库配置
// @Success 200
// @Router /api/configure/artifact/{id} [delete]
// @Param   id     path     int     true	"ImageID.ID"
// @Security JWT
func RemoveArtifactRepo(ctx *gin.Context) {
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

	err = artifact.RemoveRepo(userId, uint(repoId))

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
