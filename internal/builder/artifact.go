package builder

import (
	"encoding/json"
	"fmt"
	pipeline2 "go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func SaveDockImage(buildId uint) {
	if history, _ := repositories.GetPipelineHistory(buildId); history != nil {
		pipelineId := history.PipelineID

		if pipeline, _ := repositories.QueryPipeline(pipelineId); pipeline != nil {
			for _, step := range pipeline.PipelineSteps {
				if step.Type == pipeline2.Image {
					var script pipeline2.ArtifactScript
					json.Unmarshal([]byte(step.Script), &script)
					tag := utils.DockerImageTagBuild(buildId)
					imageAddr := fmt.Sprintf("%s/%s-%s:%s", script.Registry, pipeline.ArtifactName, history.Branch, tag)

					image := &repositories.ArtifactDockerImages{
						PipelineId: pipelineId,
						BuildId:    buildId,
						Name:       pipeline.ArtifactName,
						ArtifactRepoID: func() uint {
							if pipeline.ArtifactRepoId != nil {
								return *pipeline.ArtifactRepoId
							} else {
								return 0
							}
						}(),
						Tag:         tag,
						FullAddress: imageAddr,
					}
					_ = repositories.CreateArtifact(image)

					latestImage := *image
					latestImage.Tag = kube.LatestTag
					latestImage.FullAddress = fmt.Sprintf("%s/%s-%s:%s", script.Registry, pipeline.ArtifactName, history.Branch, latestImage.Tag)
					_ = repositories.UpsertLatestArtifact(&latestImage)
					break
				}
			}
		}
	}
}
