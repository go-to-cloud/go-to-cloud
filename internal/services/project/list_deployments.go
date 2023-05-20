package project

import (
	"encoding/json"
	"go-to-cloud/internal/models/deploy"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func deploymentMapper(repo repositories.DeploymentBase) deploy.Deployment {
	return deploy.Deployment{
		Id:           repo.ID,
		Namespace:    repo.K8sNamespace,
		K8S:          repo.K8sRepo.ID,
		Artifact:     repo.ArtifactDockerImageRepo.ArtifactRepoID,
		K8sName:      repo.K8sRepo.Name,
		ArtifactName: repo.ArtifactDockerImageRepo.Name,
		Ports: func() []struct {
			ServicePort   string `json:"servicePort"`
			ContainerPort string `json:"containerPort"`
			NodePort      string `json:"nodePort"`
		} {
			var t []struct {
				ServicePort   string `json:"servicePort"`
				ContainerPort string `json:"containerPort"`
				NodePort      string `json:"nodePort"`
			}
			if json.Unmarshal(repo.Ports, &t) != nil {
				return nil
			} else {
				return t
			}
		}(),
		Env: func() []struct {
			VarName  string `json:"text"`
			VarValue string `json:"value"`
		} {
			var t []struct {
				VarName  string `json:"text"`
				VarValue string `json:"value"`
			}
			t1 := make([]struct {
				VarName  string `json:"text"`
				VarValue string `json:"value"`
			}, 0)
			if json.Unmarshal(repo.Env, &t) != nil {
				return nil
			} else {
				for i, s := range t {
					if len(s.VarName) > 0 {
						t1 = append(t1, t[i])
					}
				}
				return t1
			}
		}(),
		Replicate:   repo.Replicas,
		CpuRequest:  repo.ResourceLimitCpuRequest,
		CpuLimits:   repo.ResourceLimitCpuLimits,
		MemRequest:  repo.ResourceLimitMemRequest,
		MemLimits:   repo.ResourceLimitMemLimits,
		ArtifactTag: repo.ArtifactTag,
		LastDeployAt: func() *utils.JsonTime {
			t := repo.LastDeployAt
			if t == nil {
				return nil
			} else {
				m := utils.JsonTime(*t)
				return &m
			}
		}(),
		Healthcheck:     repo.Liveness,
		HealthcheckPort: repo.LivenessPort,
	}
}

func ListDeployments(projectId uint) ([]deploy.Deployment, error) {
	deployments, err := repositories.QueryDeploymentsByProjectId(projectId)

	if err != nil {
		return nil, err
	}

	models := make([]deploy.Deployment, len(deployments))
	for i := range deployments {
		models[i] = deploymentMapper(deployments[i].DeploymentBase)
	}

	return models, nil
}
