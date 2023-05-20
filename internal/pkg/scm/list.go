package scm

import (
	"context"
	scm2 "github.com/drone/go-scm/scm"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/models/scm"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// List 读取仓库
// @Params:
//
//	orgs: 当前用户所在组织
//	query: 查询条件
func List(orgs []uint, query *scm.Query) ([]scm.Scm, error) {
	var orgId []uint
	if query == nil || len(query.Orgs) == 0 {
		//	默认取当前用户所属全体组织
		orgId = orgs
	} else {
		// 	计算查询条件中的所属组织与当前用户所属组织的交集
		orgId = utils.Intersect(orgs, query.Orgs)
	}

	patternName := ""
	pager := models.Pager{}
	imported := false
	if query != nil {
		patternName = query.Name
		pager = query.Pager
		imported = query.Imported
	}
	if merged, err := repositories.QueryCodeRepo(orgId, patternName, &pager, imported); err != nil {
		return nil, err
	} else {
		rlt := make([]scm.Scm, len(merged))
		for i, m := range merged {
			orgLites := make([]scm.OrgLite, len(m.Org))
			for i, lite := range m.Org {
				orgLites[i] = scm.OrgLite{
					OrgId:   lite.OrgId,
					OrgName: lite.OrgName,
				}
			}
			rlt[i] = scm.Scm{
				Testing: scm.Testing{
					Id:       m.ID,
					Origin:   scm.Type(m.ScmOrigin),
					IsPublic: m.IsPublic != 0,
					Url:      m.Url,
					Token:    &merged[i].AccessToken,
				},
				Name:     m.Name,
				OrgLites: orgLites,
				Remark:   m.Remark,
				UpdatedAt: func() string {
					if m.UpdatedAt.Before(m.CreatedAt) {
						return m.CreatedAt.Format("2006-01-02")
					} else {
						return m.UpdatedAt.Format("2006-01-02")
					}
				}(),
			}
		}
		return rlt, err
	}
}

// ListCodeProjects 列出仓库里的项目
func ListCodeProjects(origin scm.Type, url, token *string) ([]scm.RepositoryModel, error) {
	if client, err := newClient(origin, false, url, token); err != nil {
		return nil, err
	} else {
		repo, _, err := client.Repositories.List(context.Background(), scm2.ListOptions{
			Page: 0,
			Size: 100,
		})

		if err != nil {
			return nil, err
		}

		rlt := make([]scm.RepositoryModel, len(repo))
		for i, repository := range repo {
			rlt[i] = scm.RepositoryModel{
				Id:        repository.ID,
				Name:      repository.Name,
				Url:       repository.Clone,
				Namespace: repository.Namespace,
			}
		}
		return rlt, nil
	}
}
