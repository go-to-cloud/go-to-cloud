package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
	"strconv"
)

// ListBranches
// @Tags Projects
// @Description 列出仓库分支
// @Summary 列出仓库分支
// @Success 200
// @Router /api/projects/{projectId}/src/{sourceCodeId} [get]
// @Security JWT
func ListBranches(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var projectId, sourceCodeId uint64
	var err error
	projectIdStr := ctx.Param("projectId")
	if projectId, err = strconv.ParseUint(projectIdStr, 10, 64); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	sourceCodeIdStr := ctx.Param("sourceCodeId")
	if sourceCodeId, err = strconv.ParseUint(sourceCodeIdStr, 10, 64); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	if branches, err := scm.ListBranches(uint(projectId), uint(sourceCodeId)); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, project.SourceCodeBranch{
			Branches: branches,
		})
	}
	return
}
