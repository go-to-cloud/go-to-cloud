package project

import (
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/repositories"
)

func ListDeploymentHistory(projectId, deploymentId uint) ([]deploy.DeploymentHistory, error) {
	history, err := repositories.QueryDeploymentHistory(projectId, deploymentId)

	if err != nil {
		return nil, err
	}

	models := make([]deploy.DeploymentHistory, len(history))
	for i := range history {
		models[i] = deploy.DeploymentHistory{
			Deployment:   deploymentMapper(history[i].DeploymentBase),
			DeploymentId: deploymentId,
		}
	}

	return models, nil
}
