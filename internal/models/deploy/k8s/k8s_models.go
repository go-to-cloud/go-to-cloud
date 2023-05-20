package k8s

import "go-to-cloud/internal/models"

type Testing struct {
	Id         uint    `json:"id"`
	KubeConfig *string `json:"kubeconfig"`
}

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}

type K8s struct {
	Testing
	Name          string    `json:"name" form:"name"`
	Orgs          []uint    `json:"orgs" form:"orgs"`
	OrgLites      []OrgLite `json:"orgLites"`
	Remark        string    `json:"remark"`
	ServerVersion string    `json:"serverVersion"`
	UpdatedAt     string    `json:"updatedAt"`
}

type Query struct {
	models.Pager
	K8s
}
