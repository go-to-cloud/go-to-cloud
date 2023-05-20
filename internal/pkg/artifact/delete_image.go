package artifact

import (
	"errors"
	"go-to-cloud/internal/repositories"
)

// DeleteImage 删除镜像制品
func DeleteImage(userId, imageId uint) error {
	if userId == 0 || imageId == 0 {
		return errors.New("not allowed")
	}

	return repositories.DeleteImage(userId, imageId)
}

func DeleteImages(userId, artifactRepoId uint, imageId []int) error {
	if userId == 0 || len(imageId) == 0 || artifactRepoId == 0 {
		return errors.New("not allowed")
	}

	return repositories.DeleteImages(userId, artifactRepoId, imageId)
}
