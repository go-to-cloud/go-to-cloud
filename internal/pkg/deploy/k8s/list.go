package k8s

import (
	"go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
)

// List 读取仓库
// @Params:
//
//	orgs: 当前用户所在组织
//	query: 查询条件
func List(orgs []uint, query *k8s.Query) ([]k8s.K8s, error) {
	var orgId []uint
	if query == nil {
		query = &k8s.Query{}
	}
	if len(query.Orgs) == 0 {
		//	默认取当前用户所属全体组织
		orgId = orgs
	} else {
		// 	计算查询条件中的所属组织与当前用户所属组织的交集
		orgId = utils.Intersect(orgs, query.Orgs)
	}

	if merged, err := repositories.QueryK8sRepo(orgId, query.Name, &query.Pager); err != nil {
		return nil, err
	} else {
		rlt := make([]k8s.K8s, len(merged))
		var hidden = "***Hidden***"
		for i, m := range merged {
			orgLites := make([]k8s.OrgLite, len(m.Org))
			for i, lite := range m.Org {
				orgLites[i] = k8s.OrgLite{
					OrgId:   lite.OrgId,
					OrgName: lite.OrgName,
				}
			}
			rlt[i] = k8s.K8s{
				Testing: k8s.Testing{
					Id:         m.ID,
					KubeConfig: &hidden,
				},
				Name:          m.Name,
				OrgLites:      orgLites,
				Remark:        m.Remark,
				ServerVersion: m.ServerVersion,
				UpdatedAt:     m.UpdatedAt.Format("2006-01-02"),
			}
		}
		return rlt, err
	}
}
