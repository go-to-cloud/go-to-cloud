package monitor

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/monitor"
	"net/http"
)

// DeleteDeployment
// @Tags Monitor
// @Description 删除应用
// @Summary 删除应用
// @Router /api/monitor/{k8s}/apps/delete/{deploymentId} [delete]
// @Security JWT
func DeleteDeployment(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	k8sRepoId, err := getUIntParamFromQueryOrPath("k8s", ctx, false)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	deploymentId, err := getUIntParamFromQueryOrPath("deploymentId", ctx, false)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if err := monitor.DeleteDeployment(k8sRepoId, deploymentId); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx)
	}
}
