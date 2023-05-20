package repositories

import (
	"go-to-cloud/conf"
	"gorm.io/gorm/clause"
)

// DeploymentHistory k8s环境部署历史
type DeploymentHistory struct {
	DeploymentBase
	DeploymentId          uint `json:"deploymentId" gorm:"column:deployment_id;type:bigint unsigned"`
	ProjectId             uint `json:"projectId" gorm:"project_id;type:bigint unsigned;index:deployment_history_project_id"`
	K8sRepoId             uint `json:"k8sRepoId" gorm:"column:k8s_repo_id;index:deployment_history_k8s_repo_id"`
	ArtifactDockerImageId uint `json:"artifactDockerImageId" gorm:"column:artifact_docker_image_id;type:bigint unsigned;index:deployment_history_artifact_docker_image_id"`
}

func (m *DeploymentHistory) TableName() string {
	return "deployments_history"
}

func QueryDeploymentHistory(projectId, deploymentId uint) ([]DeploymentHistory, error) {
	db := conf.GetDbClient()

	var history []DeploymentHistory

	tx := db.Model(&DeploymentHistory{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ? AND deployment_id = ?", projectId, deploymentId).Order("last_deploy_at DESC")
	err := tx.Find(&history).Error

	return returnWithError(history, err)
}
