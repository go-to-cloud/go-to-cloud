package scm

import (
	"github.com/gin-gonic/gin"
	scmModels "go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
)

// Testing
// @Tags Configure
// @Description 代码仓库配置
// @Produce json
// @Accept json
// @Param   ContentBody     body     scm.Testing     true  "Request"     example(scm.Testing)
// @Security JWT
// @Success 200
// @Router /api/configure/coderepo/testing [post]
func Testing(ctx *gin.Context) {
	var req scmModels.Testing
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if success, err := scm.Ping(&req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusOK, &msg)
		return
	} else {
		response.Success(ctx, gin.H{
			"success": success,
		})
	}
}
