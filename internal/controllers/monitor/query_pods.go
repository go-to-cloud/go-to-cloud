package monitor

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/monitor"
	"net/http"
)

// QueryPods
// @Tags Monitor
// @Description 列出应用中的Pod信息
// @Success 200 {array} kube.PodDetailDescription
// @Router /api/monitor/{k8s}/pod/{deploymentId} [get]
// @Param        force    query     bool  false  "force refresh"
// @Param        deploymentId    path     string  false  "deployment id， 用于从部署方案中跳转到对应的应用"
// @Param        k8s    path     string  true  "k8s repo id"
// @Security JWT
func QueryPods(ctx *gin.Context) {
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
	deploymentId, err := getUIntParamFromQueryOrPath("deploymentId", ctx, true)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	forceRefresh := getBoolParamFromQuery("force", ctx, false)

	m, err := monitor.QueryPods(deploymentId, k8sRepoId, forceRefresh)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}
