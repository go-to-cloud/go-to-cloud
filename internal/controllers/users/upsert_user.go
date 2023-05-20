package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models/user"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/repositories"
	"net/http"
)

// UpsertUser
// @Tags User
// @Description 新建或更新用户
// @Success 200
// @Router /api/user [post]
// @Router /api/user [put]
// @Security JWT
func UpsertUser(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var err error
	var req user.User
	if err = ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if req.Id == 0 {
		err = repositories.CreateUser(&req)
	} else {
		err = repositories.UpdateUser(req.Id, &req)
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
