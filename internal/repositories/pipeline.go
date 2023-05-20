package repositories

import (
	"encoding/json"
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/pipeline"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type PipelineBase struct {
	Model
	PipelineSteps []PipelineSteps         `json:"-" gorm:"foreignKey:pipeline_id"`
	Name          string                  `json:"name" gorm:"column:name;type:nvarchar(64)"` // 计划名称
	Env           string                  `json:"env" gorm:"column:env;type:nvarchar(64)"`   // 运行环境(模板), e.g. dotnet:6; go:1.17
	SourceCode    ProjectSourceCode       `json:"-" gorm:"foreignKey:source_code_id"`
	Branch        string                  `json:"branch" gorm:"column:branch;type:nvarchar(128)"` // 分支名称
	CreatedBy     uint                    `json:"created_by" gorm:"column:created_by;type:bigint unsigned"`
	Remark        string                  `json:"remark" gorm:"column:remark;type:nvarchar(200)"`
	LastRunId     uint                    `json:"last_run_id" gorm:"column:last_run_id;type:bigint unsigned"`   // 最近一次构建记录ID，即pipeline_history.id
	LastRunAt     *time.Time              `json:"last_run_at" gorm:"column:last_run_at"`                        // 最近一次运行时间
	LastRunResult pipeline.BuildingResult `json:"last_run_result" gorm:"column:last_run_result"`                // 最近一次运行结果; 1：成功；2：取消；3：失败；0：从未执行
	ArtifactName  string                  `json:"artifact_name" gorm:"column:artifact_name;type:nvarchar(200)"` // 制品名称
}

type Pipeline struct {
	PipelineBase
	ProjectID      uint  `json:"project_id" gorm:"column:project_id;type:bigint unsigned;index:pipeline_project_id"`
	ArtifactRepoId *uint `json:"artifact_repo_id" gorm:"column:artifact_repo_id;type:bigint unsigned;index:pipeline_artifact_repo_id"`
	SourceCodeID   uint  `json:"source_code_id" gorm:"column:source_code_id;type:bigint unsigned;index:pipeline_source_code_id"`
}

func (m *Pipeline) TableName() string {
	return "pipeline"
}

// NewPlan 新建构建计划
func NewPlan(projectId uint, currentUserId uint, model *pipeline.PlanModel) (plan *Pipeline, err error) {
	steps := make(steps, 0)
	sort := 0
	// TODO: 质量检查逻辑（暂时移除）
	//err = steps.qaStep(model, &sort)
	//if err != nil {
	//	return err
	//}
	err = steps.artifactStep(model, &sort)
	if err != nil {
		return nil, err
	}

	plan = &Pipeline{
		ProjectID:      projectId,
		ArtifactRepoId: model.ArtifactRepoId,
		SourceCodeID:   model.SourceCodeID,
		PipelineBase: PipelineBase{
			Name:          model.Name,
			Env:           model.Env,
			Branch:        model.Branch,
			CreatedBy:     currentUserId,
			Remark:        model.Remark,
			ArtifactName:  model.ImageName,
			LastRunResult: 0,
			PipelineSteps: steps,
		},
	}

	tx := conf.GetDbClient()

	err = tx.Omit("updated_at").Model(&Pipeline{}).Create(plan).Error

	return
}

func QueryPipeline(id uint) (*Pipeline, error) {
	db := conf.GetDbClient()

	var p Pipeline

	tx := db.Model(&Pipeline{})

	tx = tx.Preload(clause.Associations)
	tx = tx.First(&p, id)

	return returnWithError(&p, tx.Error)
}

// QueryIncompletePipeline 查找所有进行中的任务
func QueryIncompletePipeline() ([]Pipeline, error) {
	db := conf.GetDbClient()

	var p []Pipeline

	tx := db.Model(&Pipeline{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("last_run_result = ?", pipeline.UnderBuilding).Find(&p)

	return returnWithError(p, tx.Error)
}

func QueryPipelinesByProjectId(projectId uint) ([]Pipeline, error) {
	db := conf.GetDbClient()

	var plans []Pipeline

	tx := db.Model(&Pipeline{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ?", projectId)
	err := tx.Find(&plans).Error

	return returnWithError(plans, err)
}

func QueryPipelineHistoryByProjectId(projectId, pipelineId uint) ([]PipelineHistory, error) {
	db := conf.GetDbClient()

	var plans []PipelineHistory

	tx := db.Model(&PipelineHistory{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ? AND pipeline_id = ?", projectId, pipelineId)
	err := tx.Find(&plans).Error

	return returnWithError(plans, err)
}

func DeletePlan(projectId, planId uint) error {
	db := conf.GetDbClient()

	tx := db.Model(&Pipeline{})

	tx = tx.Preload(clause.Associations)
	tx = tx.Where("project_id = ?", projectId)
	err := tx.Delete(&Pipeline{PipelineBase: PipelineBase{Model: Model{ID: planId}}}).Error

	if err == nil {
		tx = tx.Session(&gorm.Session{NewDB: true})
		tx.Model(&PipelineSteps{}).Where("ci_plan_id = ?", planId).Delete(&PipelineSteps{})
	}
	return err
}

// StartPlan 启动构建计划
func StartPlan(projectId, planId, userId uint) (*Pipeline, uint, error) {
	db := conf.GetDbClient()

	var plan Pipeline
	var historyId uint // 本次构建记录ID
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload(clause.Associations).First(&plan, planId).Error; err != nil {
			return err
		}
		var repo CodeRepo
		tx.First(&repo, plan.SourceCode.CodeRepoID)
		plan.SourceCode.CodeRepo = repo

		now := time.Now()
		state := pipeline.UnderBuilding
		history := &PipelineHistory{
			PipelineID: planId,
			Params: func() datatypes.JSON {
				j, _ := json.Marshal(plan.PipelineSteps)
				return j
			}(),
			ProjectID:    projectId,
			SourceCodeID: plan.SourceCodeID,
			PipelineBase: PipelineBase{
				Name:   plan.Name,
				Env:    plan.Env,
				Branch: plan.Branch,

				CreatedBy:     userId,
				Remark:        plan.Remark,
				LastRunAt:     &now,
				LastRunResult: state,
			},
			BuildLog: "",
		}

		if err := tx.Omit("updated_at").Create(history).Error; err != nil {
			return err
		}

		historyId = history.ID

		if err := tx.Model(&plan).Updates(&Pipeline{
			PipelineBase: PipelineBase{
				LastRunId:     historyId,
				LastRunAt:     &now,
				LastRunResult: state,
			},
		}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})

	return &plan, historyId, err
}
