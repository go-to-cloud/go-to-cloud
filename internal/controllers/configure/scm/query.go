package scm

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	scm2 "go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/pkg/scm"
	"net/http"
)

// QueryCodeRepos
// @Tags Configure
// @Description 代码仓库配置
// @Success 200 {object} scm.Scm
// @Router /api/configure/coderepo [get]
// @Security JWT
func QueryCodeRepos(ctx *gin.Context) {
	exists, _, _, orgsId, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var query scm2.Query
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Fail(ctx, http.StatusBadRequest, nil)
		return
	}

	result, err := scm.List(orgsId, &query)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, result)
}
