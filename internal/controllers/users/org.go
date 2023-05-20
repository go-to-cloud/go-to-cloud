package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/users"
	"net/http"
	"strconv"
)

// OrgList
// @Tags User
// @Description 查看用户所属组织
// @Success 200 {array} user.Org
// @Router /api/user/org/list [get]
// @Security JWT
func OrgList(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	if orgs, err := users.GetOrgList(); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, orgs)
	}
}

// Belonged
// @Tags User
// @Description 用户加入的组织
// @Success 200 {array} uint
// @Router /api/user/{userId}/belongs [get]
// @Security JWT
func Belonged(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	userIdStr := ctx.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	if o, err := users.GetUserBelongs(uint(userId)); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		id := make([]uint, len(o))
		for i, user := range o {
			id[i] = user.Id
		}
		response.Success(ctx, id)
	}
}

// Belong
// @Tags User
// @Description 将组织添加/移除用户（用户维度操作）
// @Success 200
// @Router /api/user/{userId}/join [put]
// @Param   ContentBody     body     []uint     true  "Request"     example([]uint，orgId)
// @Security JWT
func Belong(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	userIdStr := ctx.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	var tmpOrgs tmp
	if err := ctx.ShouldBindJSON(&tmpOrgs); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	if err := users.JoinOrgs(tmpOrgs.Orgs, uint(userId)); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, gin.H{
			"success": true,
		})
	}
}
