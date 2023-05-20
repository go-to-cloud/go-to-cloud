package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/repositories"
	"net/http"
	"strconv"
)

// DeleteBuildPlan 删除构建计划
// @Tags Projects
// @Description 删除构建计划
// @Summary 删除构建计划
// @Success 200
// @Router /api/projects/{projectId}/pipeline/{id} [delete]
// @Security JWT
func DeleteBuildPlan(ctx *gin.Context) {
	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	planIdStr := ctx.Param("id")
	planId, err := strconv.ParseUint(planIdStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = repositories.DeletePlan(uint(projectId), uint(planId))
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
