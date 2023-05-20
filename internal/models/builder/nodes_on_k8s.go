package builder

import "go-to-cloud/internal/models"

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}

type AvailableNodesOnK8s struct {
	Id               uint `json:"id"`
	AvailableWorkers int  `json:"availableWorkers"` // 可用节点数
}

type NodesOnK8s struct {
	AvailableNodesOnK8s
	Orgs         []uint    `json:"orgs" form:"orgs"`
	OrgLites     []OrgLite `json:"orgLites"`
	Name         string    `json:"name" form:"name"`
	Remark       string    `json:"remark"`
	KubeConfig   string    `json:"kubeConfig"`
	Workspace    string    `json:"workspace"`    // 工作空间（k8s名字空间）
	MaxWorkers   int       `json:"maxWorkers"`   // 最大同时工作数量
	AgentVersion string    `json:"agentVersion"` // 代理版本
}

type Query struct {
	models.Pager
	NodesOnK8s
}
