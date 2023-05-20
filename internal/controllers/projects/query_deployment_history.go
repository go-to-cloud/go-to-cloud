package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// QueryDeploymentHistory 获取部署方案列表
// @Tags Projects
// @Description 获取部署方案
// @Summary 获取部署方案
// @Success 200 {array} deploy.DeploymentHistory
// @Router /api/projects/{projectId}/deploy/app/{deploymentId}/history [get]
// @Security JWT
func QueryDeploymentHistory(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)
	if err != nil || projectId <= 0 {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	deploymentIdStr := ctx.Param("deploymentId")
	deploymentId, err := strconv.ParseUint(deploymentIdStr, 10, 64)
	if err != nil || deploymentId <= 0 {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	m, err := project.ListDeploymentHistory(uint(projectId), uint(deploymentId))
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}
