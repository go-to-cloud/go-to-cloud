package monitor

import (
	"go-to-cloud/internal/models/deploy"
)

func RestartApps(k8sRepoId uint, restart *deploy.RestartPods) error {

	if client, ns, name, err := getKubeClient(k8sRepoId, restart.Id); err != nil {
		return err
	} else {
		return client.Restart(&ns, &name)
	}
}
