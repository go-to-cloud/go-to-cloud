package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/repositories"
	"net/http"
)

// UpsertOrg
// @Tags User
// @Description 新建或更新组织
// @Success 200
// @Router /api/user/org [post]
// @Router /api/user/org [put]
// @Security JWT
func UpsertOrg(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var err error
	var req user.Org
	if err = ctx.ShouldBind(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if req.Id == 0 {
		err = repositories.CreateOrg(&req.Name, &req.Remark)
	} else {
		err = repositories.UpdateOrg(req.Id, &req.Name, &req.Remark)
	}
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
