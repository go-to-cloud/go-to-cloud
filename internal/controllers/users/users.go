package users

import (
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/users"
	"net/http"
	"strconv"
	"strings"
)

// Info
// @Tags User
// @Description 查看用户信息
// @Success 200
// @Router /api/user/info [get]
// @Security JWT
func Info(ctx *gin.Context) {
	exists, userId, userName, _, orgs, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	response.Success(ctx, gin.H{
		"userId":   userId,
		"userName": userName,
		"orgs":     orgs,
	})
}

// Logout
// @Tags User
// @Description 注销登录
// @Success 200
// @Router /api/user/logout [get]
// @Security JWT
func Logout(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"code": 20000,
		"data": gin.H{
			"name":   "Hello",
			"avatar": "https://i.jd.com/defaultImgs/9.jpg",
		},
	})
}

// List
// @Tags User
// @Description 列出所有用户
// @Success 200
// @Router /api/user/list [get]
// @Security JWT
func List(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	if u, err := users.GetUserList(); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, u)
	}
}

// Joined
// @Tags User
// @Description 列出加入指定组织的用户
// @Success 200
// @Router /api/user/joined/{orgId} [get]
// @Param   orgId     path     string     true  "OrgId"     example(OrgId)
// @Security JWT
func Joined(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	orgIdStr := ctx.Param("orgId")
	orgId, err := strconv.ParseUint(orgIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	if u, err := users.GetUsersByOrg(uint(orgId)); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		id := make([]uint, len(u))
		for i, user := range u {
			id[i] = user.Id
		}
		response.Success(ctx, id)
	}
}

type tmp struct {
	Users []uint `json:"users"`
	Orgs  []uint `json:"orgs"`
}

// Join
// @Tags User
// @Description 将成员加入/移除组织
// @Success 200
// @Router /api/user/join/{orgId} [put]
// @Param   orgId     path     string     true  "OrgId"     example(OrgId)
// @Param   ContentBody     body     []uint     true  "Request"     example([]uint, userId)
// @Security JWT
func Join(ctx *gin.Context) {
	exists, _, _, _, _, _ := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	orgIdStr := ctx.Param("orgId")
	orgId, err := strconv.ParseUint(orgIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	var tmpUser tmp
	if err := ctx.ShouldBindJSON(&tmpUser); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
	}

	if err := users.JoinOrg(uint(orgId), tmpUser.Users); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, gin.H{
			"success": true,
		})
	}
}

// ResetPassword
// @Tags User
// @Description 重置用户密码
// @Success 200
// @Router /api/user/{userId}/password/reset [put]
// @Param   ContentBody     body     string     true  "Request"     example(string)
// @Security JWT
func ResetPassword(ctx *gin.Context) {
	exists, currentUserId, _, _, _, kinds := utils.CurrentUser(ctx)
	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	userIdStr := ctx.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusBadRequest, &msg)
		return
	}

	var oldPassword, newPassword string
	var force bool
	if userId == 0 { // 修改自己的密码
		if !func() bool {
			for _, kind := range kinds {
				if strings.EqualFold(string(kind), string(models.Guest)) {
					return false
				}
			}
			return true
		}() {
			msg := "游客身份不允许修改密码"
			response.Fail(ctx, http.StatusForbidden, &msg)
			return
		}

		force = false
		userId = uint64(currentUserId)
		m := struct {
			OldPassword string `json:"oldPassword"`
			NewPassword string `json:"newPassword"`
		}{}

		if err := ctx.ShouldBindJSON(&m); err != nil {
			msg := err.Error()
			response.Fail(ctx, http.StatusBadRequest, &msg)
			return
		} else {
			oldPassword = m.OldPassword
			newPassword = m.NewPassword

			if len(strings.TrimSpace(newPassword)) < 6 {
				msg := "密码长度不足"
				response.Fail(ctx, http.StatusBadRequest, &msg)
				return
			}
		}
	} else { // 修改别人的密码（强制修改密码）
		if !func() bool {
			guest := false
			root := false
			for _, kind := range kinds {
				if strings.EqualFold(string(kind), string(models.Guest)) {
					guest = true
				}
				if strings.EqualFold(string(kind), string(models.Root)) {
					root = true
				}
			}
			return !guest && root
		}() {
			msg := "只允许root用户修改密码"
			response.Fail(ctx, http.StatusForbidden, &msg)
			return
		}
		force = true
	}

	if pwd, err := users.ResetPassword(uint(userId), &oldPassword, &newPassword, force); err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
	} else {
		response.Success(ctx, gin.H{
			"success":     true,
			"newPassword": *pwd,
		})
	}
}

// AllKinds
// @Tags User
// @Description 所有Kind
// @Accept json
// @Product json
// @Router /api/user/kinds [get]
// @Success 200 {array} string
func AllKinds(ctx *gin.Context) {
	response.Success(ctx, Kinds)
	return
}

var Kinds []models.KindPair

func init() {
	Kinds = []models.KindPair{
		{models.Dev, "研发", "Dev"},
		{models.Ops, "运维", "Ops"},
		{models.Guest, "游客", "Guest"},
	}
}
