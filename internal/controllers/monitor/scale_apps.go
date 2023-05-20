package monitor

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/monitor"
	"net/http"
)

// Scale
// @Tags Monitor
// @Description 伸缩容器
// @Summary 伸缩容器
// @Param   ContentBody     body     deploy.ScalePods     true  "Request"     example(deploy.ScalePods)
// @Router /api/monitor/{k8s}/apps/scale [put]
// @Security JWT
func Scale(ctx *gin.Context) {
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

	var req deploy.ScalePods
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if err := monitor.ScaleApps(k8sRepoId, &req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx)
	}
}
