package k8s

import (
	"go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// Bind 绑定k8s仓库
func Bind(model *k8s.K8s, userId uint, orgId []uint) error {
	return repositories.BindK8sRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}
