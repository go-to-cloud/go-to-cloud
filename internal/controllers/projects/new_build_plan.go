package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/repositories"
	"net/http"
	"strconv"
)

// NewBuildPlan 新建构建计划
// @Tags Projects
// @Description 新建构建计划
// @Summary 新建构建计划
// @Param   ContentBody     body     pipeline.PlanModel     true  "Request"     example(build.PlanModel)
// @Success 200
// @Router /api/projects/{projectId}/pipeline [post]
// @Security JWT
func NewBuildPlan(ctx *gin.Context) {
	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	var req pipeline.PlanModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if err := req.Valid(); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	_, err = repositories.NewPlan(uint(projectId), userId, &req)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
