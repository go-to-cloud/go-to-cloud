package artifact

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	artifactModels "go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// UpdateArtifactRepo 更新制品仓库
// @Tags Configure
// @Description 制品仓库配置
// @Success 200
// @Param   ContentBody     body     artifact.Artifact     true  "Request"     example(artifact.Artifact)
// @Router /api/configure/artifact [put]
// @Security JWT
func UpdateArtifactRepo(ctx *gin.Context) {
	var req artifactModels.Artifact
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	success, err := artifact.Ping(&req.Testing)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusOK, &msg)
		return
	}
	if !success {
		response.Fail(ctx, http.StatusOK, nil)
		return
	}

	exists, userId, _, orgs, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = artifact.Update(&req, userId, orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
