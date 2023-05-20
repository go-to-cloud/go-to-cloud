package k8s

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	k8smodel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// QueryK8sRepos
// @Tags Configure
// @Description k8s环境配置
// @Success 200 {object} k8s.K8s
// @Router /api/configure/deploy/k8s [get]
// @Security JWT
func QueryK8sRepos(ctx *gin.Context) {
	exists, _, _, orgsId, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query k8smodel.Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	result, err := k8s.List(orgsId, &query)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	// 脱敏
	for i := range result {
		*result[i].KubeConfig = "***Hidden***"
	}

	response.Success(ctx, result)
}
