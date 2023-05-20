package project

import (
	"go-to-cloud/internal/utils"
)

type SourceCodeModel struct {
	CodeRepoId uint   `json:"codeRepoId"` // 仓库ID
	Url        string `json:"url"`        // 代码地址
}

type SourceCodeImportedModel struct {
	SourceCodeModel
	CodeRepoOrigin int             `json:"codeRepoOrigin"`
	Id             uint            `json:"id"`            // 代码ID
	CreatedBy      string          `json:"createdBy"`     // 导入人
	CreatedAt      utils.JsonTime  `json:"updatedAt"`     // 导入时间
	LatestBuildAt  *utils.JsonTime `json:"latestBuildAt"` // 最近一次构建时间
}
