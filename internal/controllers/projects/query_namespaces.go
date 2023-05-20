package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
	"strconv"
)

// QueryNamespaces 获取部署环境的可用名字空间
// @Tags Projects
// @Description 根据当前用户所属组织获取部署环境的可用名字空间
// @Summary 获取部署环境的可用名字空间
// @Success 200 {array} string
// @Router /api/projects/{projectId}/deploy/{k8sRepoId}/namespaces [get]
// @Security JWT
func QueryNamespaces(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}
	k8sRepoIdStr := ctx.Param("k8sRepoId")
	k8sRepoId, err := strconv.ParseUint(k8sRepoIdStr, 10, 64)

	ns, err := project.ListNamespacesByK8sRepo(uint(k8sRepoId))
	if err == nil {
		response.Success(ctx, ns)
	} else {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	}

}
