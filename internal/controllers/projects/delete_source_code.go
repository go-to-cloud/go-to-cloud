package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// DeleteSourceCode 删除项目仓库
// @Tags Projects
// @Description 删除项目仓库
// @Success 200   {object}     response.Result{data=bool}
// @Router /api/projects/{projectId}/sourcecode/{id} [delete]
// @Param   projectId     path     int     true	"Project.ID"
// @Param   id            path     int     true	"SourceCode.ID"
// @Security JWT
func DeleteSourceCode(ctx *gin.Context) {
	sourceCodeIdStr := ctx.Param("id")
	sourceCodeId, err := strconv.ParseUint(sourceCodeIdStr, 10, 64)
	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = project.DeleteSourceCode(userId, uint(projectId), uint(sourceCodeId))

	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, nil, err)
	} else {
		response.Success(ctx, gin.H{"result": true})
	}
}
