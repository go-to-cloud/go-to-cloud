package docker_image

import (
	"go-to-cloud/internal/utils"
	"time"
)

type Image struct {
	Hash            string         `json:"hashId"`
	Name            string         `json:"name"`     // 镜像名
	FullName        string         `json:"fullName"` // 完整路径，包含项目名称，e.g. library/mysql:latest
	Tags            []Tag          `json:"tags"`     // Tag
	LatestVer       string         `json:"latestVersion"`
	LatestPublishAt time.Time      `json:"-"`
	PublishedAt     utils.JsonTime `json:"publishedAt"`
}

type Tag struct {
	ImageID     uint           `json:"imageID"`
	Tag         string         `json:"tags"`
	PublishedAt utils.JsonTime `json:"publishedAt"`
	IsLatest    bool           `json:"isLatest"`
	FullName    string         `json:"fullName"`
}
