package install

import (
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

func OnK8s(model *builder.OnK8sModel, userId uint, orgId []uint) error {
	_, err := repositories.NewBuilderNode(model, userId, utils.Intersect(model.Orgs, orgId))
	return err
}
