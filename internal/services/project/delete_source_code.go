package project

import "go-to-cloud/internal/repositories"

func DeleteSourceCode(userId, projectId, sourceCodeId uint) error {
	return repositories.DeleteProjectSourceCode(projectId, sourceCodeId)
}
