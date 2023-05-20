package project

import (
	"go-to-cloud/internal/models/project"
	"go-to-cloud/internal/repositories"
)

func List(orgId []uint) ([]project.DataModel, error) {
	proj, err := repositories.QueryProjectsByOrg(orgId)

	if err != nil {
		return nil, err
	}

	models := make([]project.DataModel, len(proj))
	for i, p := range proj {
		models[i] = project.DataModel{
			Id:     p.ID,
			Name:   p.Name,
			Remark: p.Remark,
			OrgId:  int(p.OrgId),
			Org:    p.Org.Name,
		}
	}
	return models, nil
}
