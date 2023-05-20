package scm

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	scm2 "go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
)

// BindCodeRepo 绑定代码仓库
// @Tags Configure
// @Description 代码仓库配置
// @Success 200
// @Param   ContentBody     body     scm.Scm     true  "Request"     example(scm.Scm)
// @Router /api/configure/coderepo/bind [post]
// @Security JWT
func BindCodeRepo(ctx *gin.Context) {
	var req scm2.Scm
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}
	success, err := scm.Ping(&req.Testing)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusForbidden, &msg)
		return
	}
	if !success {
		response.Fail(ctx, http.StatusForbidden, nil)
		return
	}

	exists, userId, _, orgs, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = scm.Bind(&req, userId, orgs)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
