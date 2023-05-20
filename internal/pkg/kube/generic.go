package kube

import (
	apps "k8s.io/client-go/applyconfigurations/apps/v1"
	core "k8s.io/client-go/applyconfigurations/core/v1"
)

type Kinds interface {
	core.ServiceApplyConfiguration |
		apps.DeploymentApplyConfiguration |
		core.NamespaceApplyConfiguration |
		core.PodApplyConfiguration
}
