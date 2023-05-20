package builder

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/builder"
	builder2 "go-to-cloud/internal/pkg/builder"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// QueryNodesOnK8s
// @Tags Configure
// @Description 节点管理
// @Success 200 {array} builder.NodesOnK8s
// @Router /api/configure/builder/nodes/k8s [get]
// @Security JWT
func QueryNodesOnK8s(ctx *gin.Context) {
	exists, _, _, orgsId, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query builder.Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	result, err := builder2.ListNodesOnK8s(orgsId, &query)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, result)
}

// QueryAvailableNodesOnK8s
// @Tags Configure
// @Description 查看节点可用工作单元
// @Success 200 {array} builder.NodesOnK8s
// @Router /api/configure/builder/nodes/k8s/available [get]
// @Security JWT
func QueryAvailableNodesOnK8s(ctx *gin.Context) {
	exists, _, _, orgs, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	nodes, err := builder2.ListNodesOnK8sOrderByIdle(orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	result := make(map[uint]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		result[nodes[i].NodeId] = nodes[i].Idle
	}

	response.Success(ctx, result)
}
