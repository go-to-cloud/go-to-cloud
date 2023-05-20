package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// ListImportedSourceCode
// @Tags Projects
// @Description 查看导入的代码
// @Summary 查看导入的代码
// @Success 200 {array} project.SourceCodeImportedModel
// @Router /api/projects/{projectId}/imported [get]
// @Security JWT
func ListImportedSourceCode(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)

	m, err := project.GetSourceCodeImported(uint(projectId))
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
	}
}
