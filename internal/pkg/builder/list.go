package builder

import (
	builder3 "go-to-cloud/internal/builder"
	"go-to-cloud/internal/models"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/repositories"
	"go-to-cloud/internal/utils"
	"sort"
)

type NodeStatus struct {
	NodeId uint // 节点ID
	Idle   int  // 空闲数量
}

// ListNodesOnK8sOrderByIdle 按空闲（最大可用）
func ListNodesOnK8sOrderByIdle(orgs []uint) ([]NodeStatus, error) {
	if merged, err := repositories.GetBuildNodesOnK8sByOrgId(orgs, "", nil); err != nil {
		return nil, err
	} else {
		rlt := make([]NodeStatus, len(merged))
		for i, m := range merged {
			if n, e := builder3.GetWorkingNodes(m.ID); e != nil {
				return nil, e
			} else {
				rlt[i] = NodeStatus{
					NodeId: m.ID,
					Idle:   m.MaxWorkers - n,
				}
			}
		}
		if len(rlt) > 0 {
			sort.SliceStable(rlt, func(i, j int) bool {
				return rlt[i].Idle > rlt[j].Idle
			})
		}
		return rlt, nil
	}
}

func ListNodesOnK8s(orgs []uint, query *builder.Query) ([]builder.NodesOnK8s, error) {
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
	if query != nil {
		patternName = query.Name
		pager = query.Pager
	}
	if merged, err := repositories.GetBuildNodesOnK8sByOrgId(orgId, patternName, &pager); err != nil {
		return nil, err
	} else {
		rlt := make([]builder.NodesOnK8s, len(merged))
		for i, m := range merged {
			orgLites := make([]builder.OrgLite, len(m.Org))
			for i, lite := range m.Org {
				orgLites[i] = builder.OrgLite{
					OrgId:   lite.OrgId,
					OrgName: lite.OrgName,
				}
			}
			rlt[i] = builder.NodesOnK8s{
				AvailableNodesOnK8s: builder.AvailableNodesOnK8s{
					Id: m.ID,
					//AvailableWorkers: func() int {
					//	if n, e := builder3.GetWorkingNodes(m.ID); e != nil {
					//		return 0
					//	} else {
					//		return m.MaxWorkers - n
					//	}
					//}(),
				},
				Name:         m.Name,
				OrgLites:     orgLites,
				Remark:       m.Remark,
				AgentVersion: m.AgentVersion,
				Workspace:    m.K8sWorkerSpace,
				MaxWorkers:   m.MaxWorkers,
				KubeConfig:   "***Hidden***", // func() string { return *m.DecryptKubeConfig() }(),

			}
		}
		return rlt, err
	}
}
