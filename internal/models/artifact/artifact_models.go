package artifact

type Type int

const (
	OSS Type = iota
	Docker
	Nuget
	Maven
	Npm
	S3
)

type Testing struct {
	Id         uint   `json:"id"`
	Type       Type   `json:"type"`
	IsSecurity bool   `json:"isSecurity"`
	Url        string `json:"url"`
	User       string `json:"user"`
	Password   string `json:"password"`
}

type Artifact struct {
	Testing
	Name      string    `json:"name" form:"name"`
	Orgs      []uint    `json:"orgs" form:"orgs"`
	OrgLites  []OrgLite `json:"orgLites"`
	Remark    string    `json:"remark"`
	UpdatedAt string    `json:"updatedAt"`
}

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}

type Query struct {
	Artifact
}

type FullName struct {
	Id      uint   `json:"id"` // 制品ID
	Name    string `json:"name"`
	Tag     string `json:"tag"`
	Address string `json:"address"`
}
