package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/repositories"
	"net/http"
	"strconv"
)

// DeleteOrg
// @Tags User
// @Description 删除组织
// @Success 200
// @Router /api/user/org/{orgId} [delete]
// @Security JWT
func DeleteOrg(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	var err error
	orgIdStr := ctx.Param("orgId")
	orgId, err := strconv.ParseUint(orgIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	err = repositories.DeleteOrg(uint(orgId))
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	response.Success(ctx, gin.H{
		"success": true,
	})
}
