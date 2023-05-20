package repositories

import (
	"encoding/json"
	"errors"
	"go-to-cloud/internal/models/pipeline"
	"strings"
)

type PipelineSteps struct {
	Model
	PipelinePlan   Pipeline              `json:"-" gorm:"foreignKey:pipeline_id"`
	PipelinePlanID int64                 `json:"pipeline_id" gorm:"column:pipeline_id"`
	Sort           int                   `json:"sort" gorm:"column:sort"`                    // 执行顺序
	Name           string                `json:"name" gorm:"column:name;type:nvarchar(200)"` // 步骤名称
	Script         string                `json:"script" gorm:"column:script;type:text"`      // 步骤脚本；当步骤类型为(5)部署时，script表示deployment和service的yml；为(4)生成制品时；由制品类型决定内容
	Type           pipeline.PlanStepType `json:"type" gorm:"column:type"`                    // 节点类型; 1:运行单测；2：运行lint；3：生成文档；4：生成镜像；5：部署；0：其他cli命令
}

func (m *PipelineSteps) TableName() string {
	return "pipeline_steps"
}

type steps []PipelineSteps

func (steps *steps) qaStep(model *pipeline.PlanModel, sort *int) error {
	if model.QaEnabled {
		if len(strings.TrimSpace(*model.UnitTest)) > 0 {
			*steps = append(*steps, PipelineSteps{
				Sort:   *sort,
				Name:   "单元测试",
				Script: *model.UnitTest,
				Type:   pipeline.UnitTest,
			})
			*sort++
		}
		if len(strings.TrimSpace(*model.LintCheck)) > 0 {
			*steps = append(*steps, PipelineSteps{
				Sort:   *sort,
				Name:   "Lint检查",
				Script: *model.LintCheck,
				Type:   pipeline.LintCheck,
			})
			*sort++
		}
	}
	return nil
}

func (steps *steps) artifactStep(model *pipeline.PlanModel, sort *int) error {
	if model.ArtifactEnabled {
		if url, account, password, isSecurity, origin, err := GetArtifactRepoByID(*model.ArtifactRepoId); err != nil {
			return err
		} else {
			if origin != 1 {
				return errors.New("not docker registry")
			}
			script, _ := json.Marshal(pipeline.ArtifactScript{
				Dockerfile: *model.Dockerfile,
				Context:    model.Workdir,
				Registry:   *url,
				IsSecurity: isSecurity,
				Account:    *account,
				Password:   *password,
			})
			*steps = append(*steps, PipelineSteps{
				Sort:   *sort,
				Name:   "镜像制品",
				Script: string(script),
				Type:   pipeline.Image,
			})
			*sort++
		}
	}
	return nil
}
