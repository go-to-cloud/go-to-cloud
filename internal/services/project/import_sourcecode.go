package project

import (
	"go-to-cloud/internal/models/project"
	scm2 "go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/pkg/scm"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
	"sort"
	"strings"
)

func GetCodeRepoGroupsByOrg(orgId []uint) ([]project.CodeRepoGroup, error) {

	coderepo, err := scm.List(orgId, nil)

	if err != nil {
		return nil, err
	}

	rlt := make([]project.CodeRepoGroup, len(coderepo))
	for i, s := range coderepo {
		if models, err := scm.ListCodeProjects(s.Origin, &s.Url, s.Token); err == nil {
			rlt[i].Id = s.Id
			rlt[i].Name = s.Name
			if s.Origin == scm2.Gitee {
				rlt[i].Host = "Gitee"
			} else if s.Origin == scm2.Github {
				rlt[i].Host = "Github"
			} else {
				rlt[i].Host = s.Url
			}
			rlt[i].Git = make([]project.GitSources, len(models))
			for j, model := range models {
				rlt[i].Git[j] = project.GitSources{
					GroupId:   s.Id,
					Id:        model.Id,
					Name:      model.Name,
					Url:       model.Url,
					Namespace: model.Namespace,
				}
			}
			sort.SliceStable(rlt[i].Git, func(x, y int) bool {
				return strings.Compare(rlt[i].Git[x].Url, rlt[i].Git[y].Url) == -1
			})
		} else {
			rlt[i].Id = s.Id
			rlt[i].Name = s.Name
			if s.Origin == scm2.Gitee {
				rlt[i].Host = "Gitee"
			} else if s.Origin == scm2.Github {
				rlt[i].Host = "Github"
			} else {
				rlt[i].Host = s.Url
			}
		}
	}
	return rlt, nil
}

func ImportSourceCode(projectId, codeRepoId, userId uint, req *project.SourceCodeModel) error {
	return repositories.UpsertProjectSourceCode(projectId, codeRepoId, userId, &req.Url)
}

func GetSourceCodeImported(projectId uint) ([]project.SourceCodeImportedModel, error) {
	rlt, err := repositories.QueryProjectSourceCode(projectId)
	if err != nil {
		return nil, err
	}

	models := make([]project.SourceCodeImportedModel, len(rlt))
	for i, code := range rlt {
		models[i] = project.SourceCodeImportedModel{
			SourceCodeModel: project.SourceCodeModel{
				CodeRepoId: code.CodeRepoID,
				Url:        code.GitUrl,
			},
			Id:             code.ID,
			CodeRepoOrigin: code.CodeRepo.ScmOrigin,
			CreatedBy:      code.CreatedUser.Account,
			CreatedAt:      utils.JsonTime(code.CreatedAt),
		}
	}
	return models, nil
}
