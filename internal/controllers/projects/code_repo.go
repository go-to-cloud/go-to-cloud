package projects

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// CodeRepo
// @Tags Projects
// @Description 列出当前账户已绑定的SCM平台及可见的代码仓库
// @Success 200 {array} project.CodeRepoGroup
// @Router /api/projects/coderepo [get]
// @Security JWT
func CodeRepo(ctx *gin.Context) {
	exists, _, _, orgId, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	m, err := project.GetCodeRepoGroupsByOrg(orgId)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	} else {
		response.Success(ctx, m)
		return
	}
}
