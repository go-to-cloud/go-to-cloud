package project

import (
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func ListPipelines(projectId uint) ([]pipeline.PlanCardModel, error) {
	plans, err := repositories.QueryPipelinesByProjectId(projectId)

	if err != nil {
		return nil, err
	}

	models := make([]pipeline.PlanCardModel, len(plans))
	for i, plan := range plans {
		unitTestEnabled, lintEnabled, artifactEnabled := false, false, false
		var unitTest, lint *string = nil, nil
		for _, step := range plan.PipelineSteps {
			if step.Type == 1 {
				unitTestEnabled = true
				unitTest = &step.Script
				continue
			}
			if step.Type == 2 {
				lintEnabled = true
				lint = &step.Script
				continue
			}
			if step.Type == 4 {
				artifactEnabled = true
				continue
			}
		}
		models[i] = pipeline.PlanCardModel{
			PlanModel: pipeline.PlanModel{
				Id:              plan.ID,
				Name:            plan.Name,
				Env:             plan.Env,
				SourceCodeID:    plan.SourceCodeID,
				Branch:          plan.Branch,
				QaEnabled:       unitTestEnabled || lintEnabled,
				UnitTest:        unitTest,
				LintCheck:       lint,
				ArtifactEnabled: artifactEnabled,
			},
			LastBuildAt: func() *utils.JsonTime {
				if plan.LastRunAt == nil {
					return nil
				} else {
					t := utils.JsonTime(*plan.LastRunAt)
					return &t
				}
			}(),
			LastBuildResult: plan.LastRunResult,
		}
	}

	return models, nil
}
