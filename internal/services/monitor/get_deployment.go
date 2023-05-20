package monitor

import (
	"errors"
	"fmt"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

func getKubeClient(expectedK8sRepoId, deploymentId uint) (k8sClient *kube.Client, namespace, deploymentName string, err error) {

	deployment, err := repositories.GetDeploymentById(deploymentId)

	if err != nil {
		return nil, "", "", err
	}
	if deployment == nil {
		return nil, "", "", errors.New("应用部署信息丢失")
	}

	if deployment.K8sRepoId != expectedK8sRepoId {
		return nil, "", "", errors.New("部署信息与当前环境不一致")
	}

	k8sRepo, err := repositories.QueryK8sRepoById(expectedK8sRepoId)
	if err != nil {
		return nil, "", "", err
	}
	if k8sRepo == nil {
		return nil, "", "", errors.New("部署环境丢失")
	}
	namespace = deployment.K8sNamespace
	deploymentName = fmt.Sprintf("%s-deployment", deployment.ArtifactDockerImageRepo.Name)

	k8sClient, err = kube.NewClient(&k8sRepo.KubeConfig)

	return
}
