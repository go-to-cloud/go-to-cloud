package project

import (
	"errors"
	builder2 "go-to-cloud/internal/builder"
	"go-to-cloud/internal/pkg/builder"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
)

func build(nodeId, buildId uint, plan *repositories.Pipeline) (*kube.PodSpecConfig, error) {
	if node, err := repositories.GetBuildNodesById(nodeId); err != nil {
		return nil, err
	} else {
		spec := builder2.BuildPodSpec(buildId, node, plan)

		if client, err := kube.NewClient(node.DecryptKubeConfig()); err != nil {
			return nil, err
		} else {
			err = client.Build(spec)
			if err == nil {
				builder2.ResetIdle(node)
			}
			return spec, err
		}
	}
}

func StartPipeline(userId uint, orgId []uint, projectId, pipelineId int64) error {
	plan, buildId, err := repositories.StartPlan(uint(projectId), uint(pipelineId), userId)
	if err == nil {
		if sortedIdleNodes, err := builder.ListNodesOnK8sOrderByIdle(orgId); err != nil {
			return err
		} else {
			if len(sortedIdleNodes) > 0 && sortedIdleNodes[0].Idle > 0 {
				node := sortedIdleNodes[0]
				_, err := build(node.NodeId, buildId, plan)
				return err
			} else {
				return errors.New("没有足够可运行的构建节点，请稍后再试") // TODO: 未来计划使用构建队列
			}
		}
	}
	return err
}
