package project

import (
	"encoding/json"
	"fmt"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"log"
	"strconv"
)

func StartDeploy(projectId, deployId uint) error {
	deployment, err := repositories.GetDeploymentByProjectId(projectId, deployId)
	if err != nil {
		return err
	}

	cfg := kube.AppDeployConfig{
		LabelSelector: kube.DeploymentLabelSelector,
		AppId:         strconv.Itoa(int(deployId)),
		Namespace:     deployment.K8sNamespace,
		Name:          deployment.ArtifactDockerImageRepo.Name,
		Ports: func() []kube.Port {
			var ports []struct {
				ServicePort   int `json:"servicePort,string"`
				ContainerPort int `json:"containerPort,string"`
				NodePort      int `json:"nodePort,string"`
			}
			if err := json.Unmarshal([]byte(deployment.Ports.String()), &ports); err != nil {
				return nil
			} else {
				rlt := make([]kube.Port, len(ports))
				for i, port := range ports {
					rlt[i] = kube.Port{
						ServicePort:   port.ServicePort,
						ContainerPort: port.ContainerPort,
						NodePort:      port.NodePort,
						PortName:      fmt.Sprintf("p-%d", i),
					}
				}
				return rlt
			}
		}(),
		Image: func() string {
			if u, err := repositories.GetDeploymentImageByTag(deployment.ArtifactDockerImageId, deployment.ArtifactTag); err != nil {
				log.Println(err.Error())
				return ""
			} else {
				return u
			}

		}(),
		Env: func() []kube.EnvVar {
			var env []struct {
				Name  string `json:"text"`
				Value string `json:"value"`
			}
			if json.Unmarshal([]byte(deployment.Env.String()), &env) != nil {
				return nil
			} else {
				rlt := make([]kube.EnvVar, 0)
				for _, e := range env {
					if len(e.Name) > 0 {
						rlt = append(rlt, kube.EnvVar{
							Name:  e.Name,
							Value: e.Value,
						})
					}
				}
				return rlt
			}
		}(),
		Replicas: int(deployment.Replicas),
		Liveness: func() *kube.ProbeConfigure {
			if len(deployment.Liveness) == 0 || deployment.LivenessPort == 0 {
				return nil
			} else {
				return &kube.ProbeConfigure{
					Path:             deployment.Liveness,
					Port:             int(deployment.LivenessPort),
					Delay:            3000,
					Period:           5000,
					Timeout:          3000,
					SuccessThreshold: 1,
					FailureThreshold: 3,
				}
			}
		}(),
		ResourceLimit: nil,
	}

	client, err := kube.NewClient(&deployment.K8sRepo.KubeConfig)

	if nil != err {
		return err
	}

	err = client.Launch(&cfg)
	if err != nil {
		return err
	}

	return repositories.Deployed(projectId, deployId)
}
