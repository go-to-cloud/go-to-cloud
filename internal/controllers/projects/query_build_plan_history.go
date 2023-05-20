package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// QueryBuildPlanHistory 获取构建计划历史记录
// @Tags Projects
// @Description 获取构建计划
// @Summary 获取构建计划
// @Success 200 {array} pipeline.PlanModel
// @Router /api/projects/{projectId}/pipeline/{pipelineId}/history [get]
// @Security JWT
func QueryBuildPlanHistory(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	pipelineIdStr := ctx.Param("pipelineId")
	pipelineId, err := strconv.ParseUint(pipelineIdStr, 10, 64)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	m, err := project.ListPipelineHistory(uint(projectId), uint(pipelineId))
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}
