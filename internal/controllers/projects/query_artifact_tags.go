package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// QueryArtifactTags 获取项目中的制品镜像版本列表
// @Tags Projects
// @Description 获取项目中的制品镜像
// @Summary 获取项目中的制品镜像
// @Success 200 {array} string
// @Router /api/projects/{projectId}/artifact/{artifactId}/tags [get]
// @Security JWT
func QueryArtifactTags(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)

	artifactIdStr := ctx.Param("artifactId")
	artifactId, err := strconv.ParseUint(artifactIdStr, 10, 64)

	tags, err := project.ListArtifactTagsById(uint(projectId), uint(artifactId))
	if err == nil {
		response.Success(ctx, tags)
	} else {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	}
}
