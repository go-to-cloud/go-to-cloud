package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// DeleteDeployment
// @Tags Projects
// @Description 删除部署应用
// @Summary 删除部署应用
// @Router /api/projects/{projectId}/deploy/{id} [delete]
// @Security JWT
func DeleteDeployment(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	val := ctx.Param("projectId")
	projectId, _ := strconv.ParseUint(val, 10, 64)

	deployIdStr := ctx.Param("id")
	deployId, err := strconv.ParseUint(deployIdStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if err := project.DeleteDeployment(uint(projectId), uint(deployId)); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx)
	}
}
