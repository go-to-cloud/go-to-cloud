package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	project2 "go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// ImportSourceCode
// @Tags Projects
// @Description 导入代码
// @Param   ContentBody     body     project.SourceCodeModel     true  "Request"     example(project.DataModel)
// @Success 200
// @Router /api/projects/{projectId}/import [POST]
// @Security JWT
func ImportSourceCode(ctx *gin.Context) {
	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	var req project2.SourceCodeModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = project.ImportSourceCode(uint(projectId), req.CodeRepoId, userId, &req)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
