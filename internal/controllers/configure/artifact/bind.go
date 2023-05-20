package artifact

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	artifactModels "go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// BindArtifactRepo 绑定制品仓库
// @Tags Configure
// @Description 制品仓库配置
// @Param   ContentBody     body     artifact.Artifact     true  "Request"     example(artifact.Artifact)
// @Success 200
// @Router /api/configure/artifact/bind [post]
// @Security JWT
func BindArtifactRepo(ctx *gin.Context) {
	var req artifactModels.Artifact
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	success, err := artifact.Ping(&req.Testing)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusForbidden, &msg)
		return
	}
	if !success {
		response.Fail(ctx, http.StatusForbidden, nil)
		return
	}

	exists, userId, _, orgs, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = artifact.Bind(&req, userId, orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
