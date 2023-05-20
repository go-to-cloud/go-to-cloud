package project

import (
	"go-to-cloud/internal/repositories"
)

func DeleteDeployment(projectId, id uint) error {
	return repositories.DeleteDeployment(projectId, id)
}
