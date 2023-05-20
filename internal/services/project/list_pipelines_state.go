package project

import (
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func ListPipelinesState(projectId uint) ([]pipeline.PlanCardModel, error) {
	plans, err := repositories.QueryPipelinesByProjectId(projectId)

	if err != nil {
		return nil, err
	}

	models := make([]pipeline.PlanCardModel, len(plans))

	for i, plan := range plans {
		models[i] = pipeline.PlanCardModel{
			PlanModel: pipeline.PlanModel{Id: plan.ID},
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
