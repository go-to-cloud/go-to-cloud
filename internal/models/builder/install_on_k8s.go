package builder

type OnK8sModel struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	MaxWorkers int    `json:"maxWorkers"`
	Workspace  string `json:"workspace"`
	KubeConfig string `json:"kubeConfig"`
	Orgs       []uint `json:"orgs"`
	Remark     string `json:"remark"`
}
