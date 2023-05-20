package project

import (
	"errors"
	"go-to-cloud/internal/models/project"
	"go-to-cloud/internal/repositories"
)

func UpdateProject(userId uint, model *project.DataModel) error {
	if len(model.Name) == 0 {
		return errors.New("empty name is not allowed")
	}

	return repositories.UpdateProject(model.Id, &model.Name, &model.Remark)
}
