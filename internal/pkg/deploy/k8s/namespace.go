package k8s

import (
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

func ListNamespaces(k8s *repositories.K8sRepo) ([]string, error) {
	c, e := kube.NewClient(&k8s.KubeConfig)
	if e != nil {
		return nil, e
	}

	return c.GetAllNamespaces(false)
}
