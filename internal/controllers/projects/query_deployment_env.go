package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// QueryDeploymentEnv 获取部署环境列表，目前仅支持K8s环境
// @Tags Projects
// @Description 获取部署环境列表，目前仅支持K8s环境
// @Summary 获取部署环境列表，目前仅支持K8s环境
// @Success 200
// @Router /api/projects/{projectId}/deploy/env [get]
// @Security JWT
func QueryDeploymentEnv(ctx *gin.Context) {
	exists, _, _, orgs, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	m, err := k8s.List(orgs, nil)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		rlt := make([]struct {
			Key   uint   `json:"id"`
			Value string `json:"name"`
		}, len(m))
		for i, s := range m {
			rlt[i] = struct {
				Key   uint   `json:"id"`
				Value string `json:"name"`
			}{Key: s.Id, Value: s.Name}
		}
		response.Success(ctx, rlt)
	}
}
