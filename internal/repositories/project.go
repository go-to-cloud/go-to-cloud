package repositories

import (
	"errors"
	"fmt"
	"go-to-cloud/conf"
	project2 "go-to-cloud/internal/models/project"
	"gorm.io/gorm/clause"
	"time"
)

type Project struct {
	Model
	CreatedBy uint   `json:"createdBy" gorm:"column:created_by"` // 仓库创建人
	Org       Org    `gorm:"foreignKey:org_id"`
	OrgId     uint   `json:"orgId" gorm:"column:org_id;"` // 所属组织
	Name      string `json:"name" gorm:"column:name;type:nvarchar(200)"`
	Remark    string `json:"remark" gorm:"column:remark;type:nvarchar(500)"`
}

func (m *Project) TableName() string {
	return "projects"
}

func QueryProjectsByOrg(orgs []uint) ([]Project, error) {
	db := conf.GetDbClient()

	var projects []Project

	tx := db.Model(&Project{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("org_id in ?", orgs)
	err := tx.Find(&projects).Error

	return projects, err
}
func buildProject(model *project2.DataModel, userId uint, orgId uint, gormModel *Model) (*Project, error) {
	return &Project{
		Model:     *gormModel,
		CreatedBy: userId,
		OrgId:     orgId,
		Name:      model.Name,
		Remark:    model.Remark,
	}, nil
}

func CreateProject(userId uint, orgId uint, model project2.DataModel) (uint, error) {
	g := &Model{
		CreatedAt: time.Now(),
	}

	repo, err := buildProject(&model, userId, orgId, g)
	if err != nil {
		return 0, err
	}

	tx := conf.GetDbClient()

	var total int64
	err = tx.Model(&Project{}).Where("name = ? AND org_id = ?", model.Name, orgId).Count(&total).Error
	if err != nil {
		return 0, err
	}
	if total > 0 {
		return 0, errors.New(fmt.Sprintf("project '%s' already exists", model.Name))
	}

	err = tx.Omit("updated_at").Create(&repo).Error
	return 0, err

}

func DeleteProject(userId, projectId uint) error {

	tx := conf.GetDbClient()

	// TODO: 校验当前userId是否拥有数据删除权限

	err := tx.Delete(&Project{
		Model: Model{
			ID: projectId,
		},
	}).Error

	return err
}

func UpdateProject(projectId uint, name, remark *string) error {
	tx := conf.GetDbClient()

	project := &Project{
		Model: Model{ID: projectId},
	}

	tx = tx.Model(project).Updates(Project{Name: *name, Remark: *remark})

	return tx.Error
}
