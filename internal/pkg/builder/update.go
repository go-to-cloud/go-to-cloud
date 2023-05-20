package builder

import (
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// Update 更新构建节点
func Update(model *builder.OnK8sModel, userId uint, orgId []uint) error {
	return repositories.UpdateBuilderNode(model, userId, utils.Intersect(model.Orgs, orgId))
}
