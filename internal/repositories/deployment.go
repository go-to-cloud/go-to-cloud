package repositories

import (
	"go-to-cloud/conf"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type DeploymentBase struct {
	Model
	K8sNamespace            string               `json:"k8sNamespace" gorm:"column:k8s_namespace;type:varchar(20)"`
	K8sRepo                 K8sRepo              `json:"-" gorm:"foreignKey:k8s_repo_id"`
	ArtifactDockerImageRepo ArtifactDockerImages `json:"-" gorm:"foreignKey:artifact_docker_image_id"`
	ArtifactTag             string               `json:"artifactTag" gorm:"column:artifact_tag;type:varchar(50);default:'latest';not null"` // 镜像版本
	Ports                   datatypes.JSON       `json:"ports" gorm:"column:ports"`                                                         // 端口{containerPort: 80, servicePort: 80, nodePort: 30080, portName: 'http'}
	Env                     datatypes.JSON       `json:"env" gorm:"column:env;type:text"`                                                   // 环境变量，[{'env_name': 'env_value'}]
	Replicas                uint                 `json:"replicas" gorm:"column:replicas;type:int unsigned"`                                 // 副本数量
	Liveness                string               `json:"liveness" gorm:"column:liveness;type: varchar(500)"`                                // 存活检查地址
	LivenessPort            uint                 `json:"livenessPort" gorm:"column:liveness_port;type: int unsigned"`                       // 存活检查端口
	Readiness               string               `json:"readiness" gorm:"column:readiness;type: varchar(500)"`                              // 就绪检查地址
	ReadinessPort           uint                 `json:"readinessPort" gorm:"column:readiness_port;type: int unsigned"`                     // 存活检查端口
	RollingMaxSurge         uint                 `json:"rollingMaxSurge" gorm:"column:rolling_max_surge;type:int unsigned"`                 // 滚动发布策略：激增数量上限（1～100）
	RollingMaxUnavailable   uint                 `json:"rollingMaxUnavailable" gorm:"rolling_max_unavailable;type:int unsigned"`            // 滚动发布策略：最大不可用上限(1~100)
	ResourceLimitCpuRequest uint                 `json:"resourceLimitCpuRequest" gorm:"resource_limit_cpu_request;type:int unsigned"`       // 资源限制：cpu分配数量，单位m
	ResourceLimitCpuLimits  uint                 `json:"resourceLimitCpuLimits" gorm:"resource_limit_cpu_limits;type:int unsigned"`         // 资源限制：cpu分配上限，单位m
	ResourceLimitMemRequest uint                 `json:"resourceLimitMemRequest" gorm:"resource_limit_mem_request;type:int unsigned"`       // 资源限制：内在分配数量，单位Mi
	ResourceLimitMemLimits  uint                 `json:"resourceLimitMemLimits" gorm:"resource_limit_mem_limits;type:int unsigned"`         // 资源限制：内在分配上限，单位Mi
	NodeSelector            datatypes.JSON       `json:"nodeSelector" gorm:"node_selector;"`                                                // 节点选择，json，[{"labelName": "labelValue"}]
	LastDeployAt            *time.Time           `json:"lastDeployAt" gorm:"last_deploy_at"`                                                // 最近一次部署时间
}

// Deployment k8s环境部署方案
type Deployment struct {
	DeploymentBase
	ProjectId             uint `json:"projectId" gorm:"project_id;type:bigint unsigned;index:deployment_project_id"`
	K8sRepoId             uint `json:"k8sRepoId" gorm:"column:k8s_repo_id;index:deployment_k8s_repo_id"`
	ArtifactDockerImageId uint `json:"artifactDockerImageId" gorm:"column:artifact_docker_image_id;type:bigint unsigned;index:deployment_artifact_docker_image_id"`
}

func (m *Deployment) TableName() string {
	return "deployments"
}

func QueryDeploymentsByProjectId(projectId uint) ([]Deployment, error) {
	db := conf.GetDbClient()

	var deployments []Deployment

	tx := db.Model(&Deployment{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ?", projectId)
	err := tx.Find(&deployments).Error

	return returnWithError(deployments, err)
}

func QueryDeploymentsByK8s(k8sRepoId uint) ([]Deployment, error) {
	db := conf.GetDbClient()

	var deployments []Deployment

	tx := db.Model(&Deployment{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("k8s_repo_id = ?", k8sRepoId)
	err := tx.Find(&deployments).Error

	return returnWithError(deployments, err)
}

func CreateDeployment(deployment *Deployment) (uint, error) {
	db := conf.GetDbClient()

	tx := db.Model(&Deployment{})

	err := tx.Create(deployment).Error
	if err != nil {
		return 0, err
	} else {
		return deployment.ID, err
	}
}

func DeleteDeployment(projectId, deploymentId uint) error {
	db := conf.GetDbClient()

	tx := db.Model(&Deployment{})

	return tx.Where("project_id = ?", projectId).Delete(&Deployment{}, deploymentId).Error
}

func GetDeploymentById(deploymentId uint) (*Deployment, error) {
	db := conf.GetDbClient()

	tx := db.Model(&Deployment{}).Preload(clause.Associations)
	var rlt Deployment
	err := tx.First(&rlt, deploymentId).Error
	return returnWithError(&rlt, err)
}

func GetDeploymentByProjectId(projectId, deploymentId uint) (*Deployment, error) {
	db := conf.GetDbClient()

	tx := db.Model(&Deployment{}).Preload(clause.Associations)
	var rlt Deployment
	err := tx.First(&rlt, deploymentId).Where("project_id = ?", projectId).Error
	return returnWithError(&rlt, err)
}

func Deployed(projectId, deploymentId uint) error {

	db := conf.GetDbClient()

	return db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&Deployment{})
		var rlt Deployment
		err := tx.First(&rlt, deploymentId).Where("project_id = ?", projectId).Update("last_deploy_at", time.Now()).Error
		if err != nil {
			return err
		}

		var rlt2 DeploymentHistory
		rlt2.DeploymentBase = rlt.DeploymentBase
		rlt2.ProjectId = rlt.ProjectId
		rlt2.K8sRepoId = rlt.K8sRepoId
		rlt2.ArtifactDockerImageId = rlt.ArtifactDockerImageId
		rlt2.DeploymentId = rlt.ID

		return tx.Model(&DeploymentHistory{}).Omit("id").Create(&rlt2).Error
	})
}

func UpdateDeploymentByHistory(projectId, deploymentId, historyId uint) (*Deployment, error) {

	db := conf.GetDbClient()

	var deployment Deployment
	err := db.Transaction(func(tx *gorm.DB) (e error) {
		tx.Model(&DeploymentHistory{})
		var history DeploymentHistory
		e = tx.First(&history, historyId).Where("project_id = ?", projectId).Error
		if e != nil {
			return
		}

		e = tx.First(&deployment, deploymentId).Omit("id", "created_at").Updates(history).Error
		if e != nil {
			return
		}

		e = tx.Model(&Deployment{}).Preload(clause.Associations).First(&deployment, deploymentId).Update("last_deploy_at", time.Now()).Error

		return e
	})
	return &deployment, err
}

func GetDeploymentImageByTag(imageId uint, tag string) (string, error) {
	db := conf.GetDbClient()

	var url string
	err := db.Raw(`
select a.full_address
from artifact_docker_images a
         inner join artifact_docker_images b on a.pipeline_id = b.pipeline_id
where b.id = ?
  and a.tag = ?`, imageId, tag).First(&url).Error

	return url, err
}
