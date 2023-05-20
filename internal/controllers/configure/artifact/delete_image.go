package artifact

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/controllers/utils"
	"go-to-cloud/internal/pkg/artifact"
	"go-to-cloud/internal/pkg/response"
	"net/http"
	"strconv"
	"strings"
)

// DeleteImageByHashId 按HashId删除镜像制品
// @Tags Configure
// @Description 删除镜像制品
// @Success 200
// @Router /api/configure/artifact/images/hashId/{hashId} [delete]
// @Param   hashId     path     int     true	"hashId"
// @Param   content    body     []int   true    "imagesId"
// @Security JWT
func DeleteImageByHashId(ctx *gin.Context) {
	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	val := ctx.Param("hashId")
	repoId, err := func() (uint, error) {
		m := strings.Split(val, ",")
		if len(m) != 2 {
			return 0, errors.New("incorrect hash")
		}
		m1, err := strconv.ParseUint(m[0], 10, 64)
		if err != nil {
			return 0, errors.New("incorrect hash")
		}
		return uint(m1), nil
	}()

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	var imageId []int
	if err := ctx.ShouldBind(&imageId); err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	err = artifact.DeleteImages(userId, repoId, imageId)

	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = ""
	}
	response.Success(ctx, gin.H{
		"success": err == nil,
		"message": message,
	})
}

// DeleteImage 删除镜像制品
// @Tags Configure
// @Description 删除镜像制品
// @Success 200
// @Router /api/configure/artifact/image/{imageId} [delete]
// @Param   imageId     path     int     true	"ImageID"
// @Security JWT
func DeleteImage(ctx *gin.Context) {
	val := ctx.Param("imageId")

	imageId, err := strconv.ParseUint(val, 10, 64)

	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}

	exists, userId, _, _, _, _ := utils.CurrentUser(ctx)

	if !exists {
		response.Fail(ctx, http.StatusUnauthorized, nil)
		return
	}

	err = artifact.DeleteImage(userId, uint(imageId))

	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = ""
	}
	response.Success(ctx, gin.H{
		"success": err == nil,
		"message": message,
	})
}
