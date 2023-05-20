package repositories

import (
	"errors"
	"fmt"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/kube"
	"gorm.io/gorm"
	"strings"
)

type ArtifactDockerImages struct {
	Model
	PipelineId     uint   `json:"pipelineId" gorm:"column:pipeline_id;index:artifact_docker_images_pipeline_id_index;type:bigint unsigned"`
	BuildId        uint   `json:"buildId" gorm:"column:build_id;type:bigint unsigned"`
	Name           string `json:"name" gorm:"column:name;type:varchar(200)"`
	ArtifactRepoID uint   `json:"artifactRepoId" gorm:"column:artifact_repo_id;index:artifact_docker_images_artifact_repo_id_index"`
	Tag            string `json:"tag" gorm:"column:tag;type:varchar(100)"`
	FullAddress    string `json:"fullAddress" gorm:"column:full_address;type:varchar(200)"`
}

// GetHashedCode 获取镜像唯一名称
func (m *ArtifactDockerImages) GetHashedCode() string {
	return fmt.Sprintf("%d,%s", m.ArtifactRepoID, strings.TrimRight(m.Name, ":"+m.Tag))
}

func (m *ArtifactDockerImages) TableName() string {
	return "artifact_docker_images"
}

// QueryImages 获取镜像仓库里的镜像列表
func QueryImages(artifactId uint) ([]ArtifactDockerImages, error) {
	db := conf.GetDbClient()

	var images []ArtifactDockerImages

	tx := db.Where(ArtifactDockerImages{ArtifactRepoID: artifactId})
	err := tx.Order("created_at DESC").Find(&images).Error

	return images, err
}

func QueryImagesByProject(projectId, artifactId uint) ([]ArtifactDockerImages, error) {
	sql := `select docker.*
from artifact_docker_images docker
         inner join pipeline p on docker.pipeline_id = p.id
         inner join artifact_repo ar on docker.artifact_repo_id = ar.id
where p.project_id = ?
  and ar.id = ?
  and docker.deleted_at is null
  and ar.deleted_at is null
order by docker.created_at desc`
	db := conf.GetDbClient()

	var images []ArtifactDockerImages

	tx := db.Raw(sql, projectId, artifactId).Find(&images)

	return images, tx.Error
}

func UpsertLatestArtifact(image *ArtifactDockerImages) error {
	db := conf.GetDbClient()

	return db.Transaction(func(tx *gorm.DB) error {
		var latest ArtifactDockerImages
		err := tx.Model(&ArtifactDockerImages{}).Where("pipeline_id = ? AND tag = ?", image.PipelineId, kube.LatestTag).First(&latest).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 新增
			return tx.Omit("id").Create(&image).Error
		} else if err == nil {
			// 更新
			return tx.Model(&ArtifactDockerImages{}).Where("id=?", latest.ID).Update("build_id", image.BuildId).Update("full_address", image.FullAddress).Error
		} else {
			return err
		}
	})
}

func CreateArtifact(image *ArtifactDockerImages) error {
	db := conf.GetDbClient()

	return db.Model(&ArtifactDockerImages{}).Create(image).Error
}

func DeleteImages(userId, artifactRepoId uint, imageId []int) error {
	// TODO: 校验当前userId是否拥有数据删除权限

	tx := conf.GetDbClient()

	return tx.Where("id IN ? AND artifact_repo_id = ?", imageId, artifactRepoId).Delete(&ArtifactDockerImages{}).Error
}

func DeleteImage(userId, imageId uint) error {

	tx := conf.GetDbClient()

	// TODO: 校验当前userId是否拥有数据删除权限

	err := tx.Delete(&ArtifactDockerImages{
		Model: Model{
			ID: imageId,
		},
	}).Error

	return err
}

func QueryLatestImagesByProjectId(projectId uint) ([]artifact.FullName, error) {
	db := conf.GetDbClient()

	var images []artifact.FullName

	tx := db.Raw(`
select a.id, a.name, a.tag, a.full_address address
from artifact_docker_images a
         inner join pipeline p on p.id = a.pipeline_id
         inner join (select max(d.updated_at) upd, d.pipeline_id pipelineId
                     from artifact_docker_images d
                     group by d.pipeline_id) x
                    on x.upd = a.updated_at AND x.pipelineId = a.pipeline_id
where p.project_id = ?`, projectId).Find(&images)

	return returnWithError(images, tx.Error)
}

func QueryImageTagsById(id uint) ([]string, error) {
	db := conf.GetDbClient()

	var tags []string

	tx := db.Raw(`
	select tag
		from artifact_docker_images b
		inner join (select pipeline_id
			from artifact_docker_images a
			where a.id = ?) x on x.pipeline_id = b.pipeline_id
			order by b.id desc`, id).Find(&tags)

	return returnWithError(tags, tx.Error)
}
