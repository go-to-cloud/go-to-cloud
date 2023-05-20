package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// QueryArtifacts 获取项目中的制品镜像
// @Tags Projects
// @Description 获取项目中的制品镜像
// @Summary 获取项目中的制品镜像
// @Success 200 {array} any
// @Router /api/projects/{projectId}/artifacts/{querystring} [get]
// @Security JWT
func QueryArtifacts(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	projectIdStr := ctx.Param("projectId")
	projectId, err := strconv.ParseUint(projectIdStr, 10, 64)

	queryString := ctx.Param("querystring")

	artifacts, err := project.ListArtifactsByKeywords(uint(projectId), &queryString)
	if err == nil {
		response.Success(ctx, artifacts)
	} else {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	}
}
