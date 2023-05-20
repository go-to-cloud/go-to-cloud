package scm

import (
	"errors"
	"go-to-cloud/internal/repositories"
)

// RemoveRepo 移除代码仓库
func RemoveRepo(userId, repoId uint) error {
	if userId <= 0 || repoId <= 0 {
		return errors.New("not allowed")
	}

	return repositories.DeleteCodeRepo(userId, repoId)
}
