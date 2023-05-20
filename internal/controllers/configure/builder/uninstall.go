package builder

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/builder/uninstall"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strconv"
)

// Uninstall 卸载构建节点
// @Tags Builder
// @Description 卸载构建节点
// @Success 200
// @Router /api/configure/builder/node/{id} [delete]
// @Param   id     path     int     true	"Node ID"
// @Security JWT
func Uninstall(ctx *gin.Context) {
	val := ctx.Param("id")

	nodeId, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = uninstall.OnK8s(userId, uint(nodeId))

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
