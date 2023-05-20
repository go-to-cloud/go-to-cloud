package artifact

import (
	"errors"
	"fmt"
	"go-to-cloud/internal/models/artifact"
	"go-to-cloud/internal/pkg/artifact/registry"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// List 读取制品仓库
// @Params:
//
//	orgs: 当前用户所在组织
//	query: 查询条件
func List(orgs []uint, query *artifact.Query) ([]artifact.Artifact, error) {
	var orgId []uint
	if len(query.Orgs) == 0 {
		//	默认取当前用户所属全体组织
		orgId = orgs
	} else {
		// 	计算查询条件中的所属组织与当前用户所属组织的交集
		orgId = utils.Intersect(orgs, query.Orgs)
	}

	if merged, err := repositories.QueryArtifactRepo(orgId, query.Name); err != nil {
		return nil, err
	} else {
		rlt := make([]artifact.Artifact, len(merged))
		for i, m := range merged {
			orgLites := make([]artifact.OrgLite, len(m.Org))
			for i, lite := range m.Org {
				orgLites[i] = artifact.OrgLite{
					OrgId:   lite.OrgId,
					OrgName: lite.OrgName,
				}
			}
			rlt[i] = artifact.Artifact{
				Testing: artifact.Testing{
					Id:         m.ID,
					Type:       artifact.Type(m.ArtifactOrigin),
					IsSecurity: m.IsSecurity,
					Url:        m.Url,
					User:       m.Account,
					Password:   m.Password,
				},
				Name:      m.Name,
				OrgLites:  orgLites,
				Remark:    m.Remark,
				UpdatedAt: m.UpdatedAt.Format("2006-01-02"),
			}
		}
		return rlt, err
	}
}

func ItemsList(artifactID uint) (any, error) {
	_, _, _, _, origin, err := repositories.GetArtifactRepoByID(artifactID)
	if err != nil {
		return nil, err
	}
	originType := artifact.Type(origin)
	switch originType {
	case artifact.Docker:
		return registry.QueryImages(artifactID)
	}

	return nil, errors.New(fmt.Sprintf("Not Support Origin Type Code: %d", originType))
}

// ListByProject 根据项目读取制品仓库
// @Params:
//
//	orgs: 当前用户所在组织
//	query: 查询条件
func ListByProject(projectId uint, orgs []uint) ([]artifact.Artifact, error) {
	if merged, err := repositories.QueryArtifactRepo(orgs, ""); err != nil {
		return nil, err
	} else {
		rlt := make([]artifact.Artifact, len(merged))
		for i, m := range merged {
			orgLites := make([]artifact.OrgLite, len(m.Org))
			for i, lite := range m.Org {
				orgLites[i] = artifact.OrgLite{
					OrgId:   lite.OrgId,
					OrgName: lite.OrgName,
				}
			}
			rlt[i] = artifact.Artifact{
				Testing: artifact.Testing{
					Id:         m.ID,
					Type:       artifact.Type(m.ArtifactOrigin),
					IsSecurity: m.IsSecurity,
					Url:        m.Url,
					User:       m.Account,
					Password:   m.Password,
				},
				Name:      m.Name,
				OrgLites:  orgLites,
				Remark:    m.Remark,
				UpdatedAt: m.UpdatedAt.Format("2006-01-02"),
			}
		}
		return rlt, err
	}
}

func ItemsListByProject(projectId, artifactID uint) (any, error) {
	_, _, _, _, origin, err := repositories.GetArtifactRepoByID(artifactID)
	if err != nil {
		return nil, err
	}
	originType := artifact.Type(origin)
	switch originType {
	case artifact.Docker:
		return registry.QueryImagesByProject(projectId, artifactID)
	}

	return nil, errors.New(fmt.Sprintf("Not Support Origin Type Code: %d", originType))
}
