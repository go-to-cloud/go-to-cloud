package project

import (
	"go-to-cloud/internal/pkg/deploy/k8s"
	"go-to-cloud/internal/repositories"
)

// ListNamespacesByK8sRepo 用于新建
func ListNamespacesByK8sRepo(k8sRepoId uint) ([]string, error) {
	if repo, err := repositories.QueryK8sRepoById(k8sRepoId); err != nil {
		return nil, err
	} else {
		return k8s.ListNamespaces(repo)
	}
}
