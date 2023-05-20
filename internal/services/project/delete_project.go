package project

import (
	"go-to-cloud/internal/repositories"
)

func Delete(userId, projectId uint) error {
	return repositories.DeleteProject(userId, projectId)
}
