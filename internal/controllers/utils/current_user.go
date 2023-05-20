package utils

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go-to-cloud/internal/models"
	"strconv"
)

func CurrentUser(ctx *gin.Context) (exists bool, userId uint, user *string, orgIds []uint, orgs map[uint]string, kinds []models.Kind) {

	defer func() {
		if r := recover(); r != nil {
			exists = false
		}
	}()

	mapping := ctx.MustGet("JWT_PAYLOAD").(jwt.MapClaims)

	jti := mapping["jti"].(float64)
	sub := mapping["sub"].(string)

	orgsMaps := mapping["orgs"]
	if orgsMaps != nil {
		maps := orgsMaps.(map[string]interface{})
		if sz := len(maps); sz > 0 {
			orgs = make(map[uint]string, sz)
			orgIds = make([]uint, 0)
			for key, val := range maps {
				orgId, _ := strconv.ParseUint(key, 10, 64)
				orgs[uint(orgId)] = val.(string)
				orgIds = append(orgIds, uint(orgId))
			}
		}
	}

	userId = uint(jti)
	user = &sub

	kindsStr := mapping["kind"].([]interface{})
	for _, i2 := range kindsStr {
		kinds = append(kinds, models.Kind(i2.(string)))
	}

	exists = true

	return
}
