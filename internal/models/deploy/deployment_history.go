package deploy

type DeploymentHistory struct {
	Deployment
	DeploymentId uint `json:"deploymentId"`
}
