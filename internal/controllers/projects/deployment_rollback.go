package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// Rollback
// @Tags Projects
// @Description 回滚部署
// @Summary 回滚部署
// @Router /api/projects/{projectId}/deploy/{id}/rollback/{historyId} [put]
// @Security JWT
func Rollback(ctx *gin.Context) {
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

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	historyIdStr := ctx.Param("historyId")
	historyId, err := strconv.ParseUint(historyIdStr, 10, 64)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	if err := project.RollbackDeploy(uint(projectId), uint(id), uint(historyId)); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx)
	}
}
