package monitor

import (
	"errors"
	"go-to-cloud/internal/repositories"
)

func DeleteDeployment(expectedK8sRepoId, deploymentId uint) error {

	deployment, err := repositories.GetDeploymentById(deploymentId)
	if err != nil {
		return err
	}
	if deployment == nil {
		return errors.New("应用部署信息丢失")
	}

	if deployment.K8sRepoId != expectedK8sRepoId {
		return errors.New("部署信息与当前环境不一致")
	}

	k8sRepo, err := repositories.QueryK8sRepoById(expectedK8sRepoId)
	if err != nil {
		return err
	}
	if k8sRepo == nil {
		return errors.New("部署环境丢失")
	}

	return repositories.DeleteDeployment(deployment.ProjectId, deploymentId)
}
