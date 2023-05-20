package k8s

import (
	"github.com/gin-gonic/gin"
	k8sModel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// Testing
// @Tags Configure
// @Description k8s仓库配置
// @Produce json
// @Accept json
// @Param   ContentBody     body     k8s.Testing     true  "Request"     example(k8s.Testing)
// @Security JWT
// @Success 200
// @Router /api/configure/deploy/k8s/testing [post]
func Testing(ctx *gin.Context) {
	var req k8sModel.Testing
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if success, err := k8s.Ping(&req); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusOK, &msg)
		return
	} else {
		response.Success(ctx, gin.H{
			"success": success,
		})
	}
}
