package artifact

import (
	"errors"
	"go-to-cloud/internal/repositories"
)

// RemoveRepo 移除制品仓库
func RemoveRepo(userId, repoId uint) error {
	if userId <= 0 || repoId <= 0 {
		return errors.New("not allowed")
	}

	return repositories.DeleteArtifactRepo(userId, repoId)
}
