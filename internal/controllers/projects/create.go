package projects

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	project2 "go-to-cloud/internal/models/project"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/project"
	"net/http"
)

// Create
// @Tags Projects
// @Description 创建新的项目
// @Param   ContentBody     body     project.DataModel     true  "Request"     example(project.DataModel)
// @Success 200
// @Router /api/projects [POST]
// @Security JWT
func Create(ctx *gin.Context) {
	var req project2.DataModel
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, err)
		return
	}

	if req.OrgId < 0 {
		response.BadRequest(ctx, errors.New("one organization at least"))
		return
	}

	exists, userId, _, orgs, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	_, err := project.CreateNewProject(userId, orgs, req)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
