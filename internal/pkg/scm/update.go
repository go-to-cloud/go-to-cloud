package scm

import (
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// Update 更新代码仓库
func Update(model *scm.Scm, userId uint, orgId []uint) error {
	return repositories.UpdateCodeRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}
