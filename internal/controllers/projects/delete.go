package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// DeleteProject 删除项目仓库
// @Tags Projects
// @Description 删除项目仓库
// @Success 200
// @Router /api/projects/{projectId} [delete]
// @Param   id     path     int     true	"Project.ID"
// @Security JWT
func DeleteProject(ctx *gin.Context) {
	val := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = project.Delete(userId, uint(projectId))

	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, nil, err)
	} else {
		response.Success(ctx)
	}
}
