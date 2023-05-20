package monitor

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/monitor"
	"net/http"
)

// DeletePod
// @Tags Monitor
// @Description 删除容器
// @Summary 删除容器
// @Param   ContentBody     body     deploy.DeletePod     true  "Request"     example(deploy.DeletePod)
// @Router /api/monitor/{k8s}/apps/delete [put]
// @Security JWT
func DeletePod(ctx *gin.Context) {
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

	var req deploy.DeletePod
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if err := monitor.DeletePod(k8sRepoId, &req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx)
	}
}
