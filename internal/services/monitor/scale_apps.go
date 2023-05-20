package monitor

import (
	"go-to-cloud/internal/models/deploy"
)

func ScaleApps(k8sRepoId uint, scale *deploy.ScalePods) error {
	if client, ns, name, err := getKubeClient(k8sRepoId, scale.Id); err != nil {
		return err
	} else {
		return client.Scale(&ns, &name, scale.Num)
	}
}
