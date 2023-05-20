package project

import "go-to-cloud/internal/repositories"

// ListArtifactTagsById 根据制品ID查询镜像历史版本（tags)
func ListArtifactTagsById(projectId uint, artifactId uint) ([]string, error) {
	return repositories.QueryImageTagsById(artifactId)
}
