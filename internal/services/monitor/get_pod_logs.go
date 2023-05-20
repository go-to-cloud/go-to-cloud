package monitor

import (
	"context"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

func followLogs(ctx context.Context, k8sId, deploymentId uint, podName, containerName string, previous bool, logs func([]byte)) error {
	repo, err := repositories.QueryK8sRepoById(k8sId)
	if err != nil {
		return err
	}

	deployment, err := repositories.GetDeploymentById(deploymentId)
	if err != nil {
		return err
	}

	client, err := kube.NewClient(&repo.KubeConfig)
	if err != nil {
		return err
	}

	lines := int64(100)
	if len(containerName) == 0 {
		containerName = deployment.ArtifactDockerImageRepo.Name
	}
	stream, err := client.GetPodStreamLogs(ctx, deployment.K8sNamespace, podName, containerName, &lines, true, previous)
	if err != nil {
		return err
	}

	defer stream.Close()

	p := make([]byte, 1024)
	for {
		n, err := stream.Read(p)
		if err != nil {
			return err
		}

		logs(p[:n])
	}
}
