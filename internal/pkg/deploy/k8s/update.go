package k8s

import (
	k8sModel "go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// Update 更新代码仓库
func Update(model *k8sModel.K8s, userId uint, orgId []uint) error {
	return repositories.UpdateK8sRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}
