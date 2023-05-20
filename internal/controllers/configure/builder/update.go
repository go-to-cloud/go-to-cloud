package builder

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/builder"
	builder2 "go-to-cloud/internal/pkg/builder"
	"go-to-cloud/internal/pkg/response"
	"net/http"
)

// UpdateBuilderNode 更新构建节点信息
// @Tags Configure
// @Description 更新构建节点信息
// @Success 200
// @Param   ContentBody     body     builder.OnK8sModel     true  "Request"     example(builder.OnK8sModel)
// @Router /api/configure/builder/node [put]
// @Security JWT
func UpdateBuilderNode(ctx *gin.Context) {
	var req builder.OnK8sModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	exists, userId, _, orgs, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	if err := builder2.Update(&req, userId, orgs); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, gin.H{
			"success": true,
		})
	}
}
