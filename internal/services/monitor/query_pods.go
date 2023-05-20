package monitor

import (
	"context"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

// QueryPods 查找Pods
func QueryPods(deploymentId, k8sId uint, force bool) ([]kube.PodDetailDescription, error) {
	repo, err := repositories.QueryK8sRepoById(k8sId)
	if err != nil {
		return nil, err
	}

	client, err := kube.NewClient(&repo.KubeConfig)
	if err != nil {
		return nil, err
	}

	deployment, err := repositories.GetDeploymentById(deploymentId)
	if err != nil {
		return nil, err
	}

	if !force {

	}

	pods, err := client.GetPods(context.TODO(), deployment.K8sNamespace, "", func() string {
		return "app=" + deployment.ArtifactDockerImageRepo.Name
	}, force)
	if err != nil {
		return nil, err
	}

	return pods, nil
}
