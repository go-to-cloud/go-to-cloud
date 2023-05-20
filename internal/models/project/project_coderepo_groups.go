package project

type GitSources struct {
	GroupId   uint   `json:"groupId"`
	Id        string `json:"id"`
	Name      string `json:"label"`
	Namespace string `json:"namespace"`
	Url       string `json:"value"`
}
type CodeRepoGroup struct {
	Id   uint         `json:"id"`
	Name string       `json:"label"`
	Host string       `json:"host"`
	Git  []GitSources `json:"options"`
}
