package builder

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/builder"
	k8sModel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/builder/install"
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// K8sInstall 安装构建节点(基于k8s)
// @Tags Builder
// @Description 安装构建节点
// @Success 200
// @Router /api/configure/builder/install/k8s [post]
// @Param   ContentBody     body     builder.OnK8sModel     true  "Request"     example(builder.OnK8sModel)
// @Security JWT
func K8sInstall(ctx *gin.Context) {
	var req builder.OnK8sModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if len(req.Workspace) == 0 {
		req.Workspace = "gotocloud-agent"
	}

	exists, userId, _, orgs, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	if _, err := k8s.Ping(&k8sModel.Testing{
		KubeConfig: &req.KubeConfig,
	}); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusForbidden, &msg)
		return
	}

	if err := install.OnK8s(&req, userId, orgs); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
