package project

import (
	"encoding/json"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/repositories"
	"gorm.io/datatypes"
)

func CreateDeployments(projectId uint, d *deploy.Deployment) (uint, error) {
	ser := func(v any) string {
		if m, e := json.Marshal(v); e != nil {
			return ""
		} else {
			return string(m)
		}
	}
	v := func(l uint) uint {
		if d.EnableLimit {
			return l
		} else {
			return 0
		}
	}
	repo := repositories.Deployment{
		ProjectId:             projectId,
		K8sRepoId:             d.K8S,
		ArtifactDockerImageId: d.Artifact,
		DeploymentBase: repositories.DeploymentBase{
			K8sNamespace:            d.Namespace,
			ArtifactTag:             d.ArtifactTag,
			Ports:                   datatypes.JSON(ser(d.Ports)),
			Env:                     datatypes.JSON(ser(d.Env)),
			Replicas:                d.Replicate,
			ResourceLimitCpuRequest: v(d.CpuRequest),
			ResourceLimitCpuLimits:  v(d.CpuLimits),
			ResourceLimitMemRequest: v(d.MemRequest),
			ResourceLimitMemLimits:  v(d.MemLimits),
			Liveness:                d.Healthcheck,
			LivenessPort:            d.HealthcheckPort,
		},
	}
	return repositories.CreateDeployment(&repo)
}
