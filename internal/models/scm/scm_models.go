package scm

import (
	"github.com/drone/go-scm/scm"
	"go-to-cloud/internal/models"
)

type Type int

const (
	Gitlab Type = iota
	Github
	Gitee
	Gitea
)

type Testing struct {
	Id       uint    `json:"id"`
	Origin   Type    `json:"origin"`
	IsPublic bool    `json:"isPublic"`
	Url      string  `json:"url"`
	Token    *string `json:"token"`
}

type OrgLite struct {
	OrgId   uint   `json:"orgId"`
	OrgName string `json:"orgName"`
}

type Scm struct {
	Testing
	Name      string    `json:"name" form:"name"`
	Orgs      []uint    `json:"orgs" form:"orgs"`
	OrgLites  []OrgLite `json:"orgLites"`
	Remark    string    `json:"remark"`
	UpdatedAt string    `json:"updatedAt"`
}

type Branch struct {
	*scm.Reference
}

type Query struct {
	models.Pager
	Scm
	Imported bool `json:"-"` // 是否已被导入项目
}
