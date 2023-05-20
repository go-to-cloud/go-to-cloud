package users

import (
	"github.com/patrickmn/go-cache"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/repositories"
	"time"
)

const keyPrefix = "resource_rules_key_"

// GetAuthCodes 获取用户权限点
func GetAuthCodes(kinds []models.Kind) []models.AuthCode {

	orderedKinds := map[models.Kind]int{
		models.Guest: 0,
		models.Dev:   10,
		models.Ops:   20,
		models.Root:  99,
	}
	kindId := 0
	kind := models.Guest
	for i, m := range kinds {
		if a, ok := orderedKinds[m]; ok {
			if a > kindId {
				kindId = a
				kind = kinds[i]
			}
		}
	}

	cacheKey := keyPrefix + string(kind)
	if v, ok := resourceRules.Get(cacheKey); ok {
		return v.([]models.AuthCode)
	}

	rules := repositories.GetResourceRules()
	m := make([]models.AuthCode, 0, len(rules))
	for _, rule := range rules {
		if orderedKinds[kind] >= orderedKinds[rule.Kind] {
			m = append(m, rule.AuthCode)
		}
	}

	go resourceRules.Set(cacheKey, m, 0)
	return m
}

var resourceRules *cache.Cache

func init() {
	resourceRules = cache.New(time.Minute, 0)
}
