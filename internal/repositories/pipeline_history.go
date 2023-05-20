package repositories

import (
	"go-to-cloud/conf"
	"go-to-cloud/internal/models/pipeline"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

//type PipelineHistory struct {
//	Model
//	PipelineSteps  []PipelineSteps         `json:"-" gorm:"foreignKey:pipeline_id"`
//	ProjectID      uint                    `json:"project_id" gorm:"column:project_id;index:pipeline_history_project_id_index"`
//	Name           string                  `json:"name" gorm:"column:name;type:nvarchar(64)"`     // 计划名称
//	Env            string                  `json:"env" gorm:"column:env;type:nvarchar(64)"`       // 运行环境(模板), e.g. dotnet:6; go:1.17
//	SourceCodeID   uint                    `json:"source_code_id" gorm:"column:source_code_id"`   // 代码仓库ID
//	Branch         string                  `json:"branch" gorm:"column:branch;type:nvarchar(64)"` // 分支名称
//	BranchCommitId string                  `json:"branch_commit_id"  gorm:"column:branch_commit_id;type:nvarchar(100)"`
//	Params         datatypes.JSON          `json:"params" gorm:"column:params"`                    // 本次运行的参数(json格式）
//	CreatedBy      uint                    `json:"created_by" gorm:"column:created_by"`            // 构建人
//	Remark         string                  `json:"remark" gorm:"column:remark"`                    // 备注
//	BuildAt        time.Time               `json:"build_at" gorm:"column:build_at"`                // 运行时间
//	BuildResult    pipeline.BuildingResult `json:"run_result" gorm:"column:run_result;default:99"` // 运行结果; 1：成功；2：取消；3：失败；0：从未执行; 99：正在构建
//	BuildLog       string                  `json:"log" gorm:"column:log;type:text"`                // 构建日志
//	PipelineID     uint                    `json:"pipeline_id" gorm:"column:pipeline_id;index:pipeline_history_pipeline_id_index"`
//}

type PipelineHistory struct {
	PipelineBase
	ProjectID      uint           `json:"project_id" gorm:"column:project_id;type:bigint unsigned;index:pipeline_history_project_id"`
	ArtifactRepoId *uint          `json:"artifact_repo_id" gorm:"column:artifact_repo_id;type:bigint unsigned;index:pipeline_history_artifact_repo_id"`
	SourceCodeID   uint           `json:"source_code_id" gorm:"column:source_code_id;type:bigint unsigned;index:pipeline_history_source_code_id"`
	Params         datatypes.JSON `json:"params" gorm:"column:params"`           // 本次运行的参数(json格式）
	BuildLog       string         `json:"log" gorm:"column:log;type:mediumtext"` // 构建日志
	PipelineID     uint           `json:"pipeline_id" gorm:"column:pipeline_id;index:pipeline_history_pipeline_id_index"`
}

func (m *PipelineHistory) TableName() string {
	return "pipeline_history"
}

func GetPipelineHistory(historyId uint) (*PipelineHistory, error) {
	db := conf.GetDbClient()

	var rlt PipelineHistory
	tx := db.Model(&PipelineHistory{}).First(&rlt, historyId)

	return returnWithError(&rlt, tx.Error)
}

func UpdatePipeline(historyId uint, rlt pipeline.BuildingResult, buildLog *string) error {
	db := conf.GetDbClient()

	return db.Transaction(func(tx *gorm.DB) (err error) {

		var history PipelineHistory
		err = tx.Model(&PipelineHistory{}).First(&history, historyId).Error
		if err != nil {
			return
		}

		if history.LastRunResult == rlt {
			return
		}

		err = tx.Model(&PipelineHistory{}).Where("id = ? AND last_run_result != ?", historyId, int(rlt)).Update("last_run_result", int(rlt)).Error
		if err != nil {
			return
		}

		if len(history.BuildLog) == 0 {
			if pipeline.IsComplete(rlt) {
				log := func() *string {
					if buildLog == nil {
						empty := ""
						return &empty
					} else {
						return buildLog
					}
				}()
				err = tx.Model(&PipelineHistory{}).Where("id = ?", history.ID).Update("log", *log).Error
			}
			if err != nil {
				return
			}
		}

		tx = tx.Session(&gorm.Session{NewDB: true})
		err = tx.Model(&Pipeline{}).Where("id = ? AND last_run_id = ? AND last_run_result != ?", history.PipelineID, history.ID, int(rlt)).Update("last_run_result", int(rlt)).Error

		// 返回 nil 提交事务
		return
	})
}
