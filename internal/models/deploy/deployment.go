package deploy

import "go-to-cloud/internal/utils"

type Deployment struct {
	Id              uint   `json:"id"`
	K8S             uint   `json:"k8s"`
	K8sName         string `json:"k8sName"`
	Namespace       string `json:"namespace"`
	Artifact        uint   `json:"artifact"`
	ArtifactName    string `json:"artifactName"`
	ArtifactTag     string `json:"artifactTag"`
	Replicate       uint   `json:"replicate"`
	Healthcheck     string `json:"healthcheck"`
	HealthcheckPort uint   `json:"healthcheckPort,string"`
	EnableLimit     bool   `json:"enableLimit"`
	CpuLimits       uint   `json:"cpuLimits,string"`
	CpuRequest      uint   `json:"cpuRequest,string"`
	MemLimits       uint   `json:"memLimits,string"`
	MemRequest      uint   `json:"memRequest,string"`
	Ports           []struct {
		ServicePort   string `json:"servicePort"`
		ContainerPort string `json:"containerPort"`
		NodePort      string `json:"nodePort"`
	} `json:"ports"`
	Env []struct {
		VarName  string `json:"text"`
		VarValue string `json:"value"`
	} `json:"env"`
	LastDeployAt *utils.JsonTime `json:"lastDeployAt"`
}

type Base struct {
	Id uint `json:"id,omitempty"` // deployment id
}

type ScalePods struct {
	Base
	Num uint `json:"num,omitempty" json:"num,omitempty"`
}

type RestartPods struct {
	Base
}

type DeletePod struct {
	Base
	PodName string `json:"podName"`
}
