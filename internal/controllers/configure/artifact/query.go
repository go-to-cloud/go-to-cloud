package artifact

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	artifactModels "go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strconv"
)

// QueryArtifactRepo
// @Tags Configure
// @Description 制品仓库配置
// @Success 200 {array} artifact.Artifact
// @Router /api/configure/artifact [get]
// @Security JWT
func QueryArtifactRepo(ctx *gin.Context) {
	exists, _, _, orgsId, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query artifactModels.Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	result, err := artifact.List(orgsId, &query)

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

// QueryArtifactItems 获取仓库里的制品
// @Tags Configure
// @Description 制品仓库配置
// @Success 200 {array} artifact.Artifact
// @Router /api/configure/artifact/{id} [get]
// @Param   id     path     int     true	"ImageID.ID"
// @Security JWT
func QueryArtifactItems(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	val := ctx.Param("id")

	repoId, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	result, err := artifact.ItemsList(uint(repoId))

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, result)
}
