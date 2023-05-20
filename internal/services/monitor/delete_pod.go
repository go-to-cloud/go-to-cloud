package monitor

import (
	"context"
	"go-to-cloud/internal/models/deploy"
)

func DeletePod(k8sRepoId uint, del *deploy.DeletePod) error {
	if client, ns, _, err := getKubeClient(k8sRepoId, del.Id); err != nil {
		return err
	} else {
		return client.DeletePod(context.TODO(), ns, del.PodName)
	}
}
