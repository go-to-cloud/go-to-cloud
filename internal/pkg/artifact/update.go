package artifact

import (
	"go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// Update 更新制品仓库
func Update(model *artifact.Artifact, userId uint, orgId []uint) error {
	return repositories.UpdateArtifactRepo(model, userId, utils.Intersect(model.Orgs, orgId))
}
