package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// QueryBuildPlanState 获取构建计划状态（流水线状态）
// @Tags Projects
// @Description 获取构建计划状态
// @Summary 获取构建计划状态
// @Success 200 {array} pipeline.PlanCardModel
// @Router /api/projects/{projectId}/pipeline/state [get]
// @Security JWT
func QueryBuildPlanState(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)

	m, err := project.ListPipelinesState(uint(projectId))
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}
