package projects

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// StartBuildPlan 启动构建计划
// @Tags Projects
// @Description 启动构建计划
// @Success 200
// @Router /api/projects/{projectId}/pipeline/{id}/build [POST]
// @Security JWT
func StartBuildPlan(ctx *gin.Context) {
	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseInt(projectIdStr, 10, 64)
	if err != nil || projectId <= 0 {
		response.BadRequest(ctx, errors.New("project id not found"))
		return
	}

	pipelineIdStr := ctx.Param("id")
	pipelineId, err := strconv.ParseInt(pipelineIdStr, 10, 64)
	if err != nil || pipelineId <= 0 {
		response.BadRequest(ctx, errors.New("pipeline id not found"))
		return
	}

	exists, userId, _, orgId, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = project.StartPipeline(userId, orgId, projectId, pipelineId)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
