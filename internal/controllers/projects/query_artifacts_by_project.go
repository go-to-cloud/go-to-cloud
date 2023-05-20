package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	artifact2 "go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strconv"
)

// QueryArtifactItemsByProjectId 获取项目制品仓库中的制品镜像
// @Tags Projects
// @Description 获取项目中的制品镜像
// @Summary 获取项目中的制品镜像
// @Success 200 {array} docker_image.Image
// @Router /api/projects/{projectId}/artifact/{artifactId} [get]
// @Param   projectId     path     int     true	"Project ID"
// @Param   artifactId     path     int     true	"Artifact Repo ID"
// @Security JWT
func QueryArtifactItemsByProjectId(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	artifactIdStr := ctx.Param("artifactId")
	artifactId, err := strconv.ParseUint(artifactIdStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	result, err := artifact.ItemsListByProject(uint(projectId), uint(artifactId))

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, result)
}

// QueryArtifactsByProjectId 按项目获取制品仓库
// @Tags Projects
// @Description 按项目获取仓库里的制品
// @Success 200 {array} artifact.Artifact
// @Router /api/projects/{projectId}/artifact [get]
// @Param   projectId     path     int     true	"Project ID"
// @Security JWT
func QueryArtifactsByProjectId(ctx *gin.Context) {
	exists, _, _, orgsId, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	result, err := artifact.List(orgsId, &artifact2.Query{})

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	for i := range result {
		result[i].Password = ""
	}
	response.Success(ctx, result)
}
