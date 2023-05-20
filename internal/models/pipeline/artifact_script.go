package pipeline

type ArtifactScript struct {
	Dockerfile string `json:"dockerfile"`
	Context    string `json:"context"`
	Registry   string `json:"registry"`
	IsSecurity bool   `json:"isSecurity"`
	Account    string `json:"account"`
	Password   string `json:"password"`
}
