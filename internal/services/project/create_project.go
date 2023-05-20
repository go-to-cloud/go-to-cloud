package project

import (
	"errors"
	project2 "go-to-cloud/internal/models/project"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// CreateNewProject 创建新项目
func CreateNewProject(userId uint, orgs []uint, model project2.DataModel) (uint, error) {
	orgId := utils.Intersect(orgs, []uint{uint(model.OrgId)})
	if len(orgId) != 1 {
		return 0, errors.New("invalid organization")
	}
	return repositories.CreateProject(userId, orgId[0], model)
}
