package k8s

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	k8sModel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// BindK8sRepo 绑定K8s仓库
// @Tags Configure
// @Description k8s仓库配置
// @Success 200
// @Param   ContentBody     body     k8s.K8s     true  "Request"     example(k8s.K8s)
// @Router /api/configure/deploy/k8s/bind [post]
// @Security JWT
func BindK8sRepo(ctx *gin.Context) {
	var req k8sModel.K8s
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	ver, err := k8s.Ping(&req.Testing)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusForbidden, &msg)
		return
	}
	if len(ver) == 0 {
		response.Fail(ctx, http.StatusForbidden, nil)
		return
	} else {
		req.ServerVersion = ver
	}

	exists, userId, _, orgs, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = k8s.Bind(&req, userId, orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
