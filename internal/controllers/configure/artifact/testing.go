package artifact

import (
	"github.com/gin-gonic/gin"
	artifactModels "go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// Testing
// @Tags Configure
// @Description 制品仓库配置
// @Produce json
// @Accept json
// @Param   ContentBody     body     artifact.Testing     true  "Request"     example(artifact.Testing)
// @Security JWT
// @Success 200
// @Router /api/configure/artifact/testing [post]
func Testing(ctx *gin.Context) {
	var req artifactModels.Testing
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if success, err := artifact.Ping(&req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusOK, &msg)
		return
	} else {
		response.Success(ctx, gin.H{
			"success": success,
		})
	}
}
