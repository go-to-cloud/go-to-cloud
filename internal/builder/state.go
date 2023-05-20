package builder

import (
	"context"
	"fmt"
	"go-to-cloud/internal/models/pipeline"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"sync"
	"time"
)

type State struct {
}

// getAndSetNodesState 获取并更新所有构建节点信息
func getAndSetNodesState() {
	getAndSetK8sNodesState()
}

var allK8sPipelines map[uint]*kube.PodDescription

var artifactWatcher chan uint // pipeline_history.ID

func init() {
	allK8sPipelines = make(map[uint]*kube.PodDescription)
	artifactWatcher = make(chan uint, 5)
}

// PipelinesWatcher 流水线监测
func PipelinesWatcher() {
	// 定时获取构建节点状态
	go func() {
		for {
			c := time.Tick(time.Second * 15)
			<-c
			getAndSetNodesState()
		}
	}()

	// 制品生成监控
	go func() {
		for builderId := range artifactWatcher {
			go SaveDockImage(builderId)
		}
	}()
}

func getAndSetK8sNodesState() {
	var lock sync.Mutex
	ctx := context.TODO()

	currentExistsPipeline := make(map[uint]bool) // 当前存在于pod中的流水线ID

	if nodes, err := repositories.GetBuildNodesOnK8sByOrgId(nil, "", nil); err == nil {
		for _, node := range nodes {
			if client, err := kube.NewClient(node.DecryptKubeConfig()); err == nil {
				if pods, err := client.GetPods(ctx, node.K8sWorkerSpace, BuildIdSelectorLabel, func() string {
					return "builder=" + NodeSelectorLabel
				}, true); err == nil {
					for i, pod := range pods {
						lock.Lock()
						allK8sPipelines[pod.BuildId] = &pods[i].PodDescription
						currentExistsPipeline[pod.BuildId] = true
						lock.Unlock()

						rlt := func() pipeline.BuildingResult {
							switch pods[i].Status {
							case string(kube.Pending), string(kube.Running):
								return pipeline.UnderBuilding
							case string(kube.Succeeded):
								// TODO: 根据容器的状态最终确定构建结果
								return pipeline.BuildingSuccess
							case string(kube.Failed):
								return pipeline.BuildingFailed
							default:
								return pipeline.NeverBuild
							}
						}()

						if pipeline.IsComplete(rlt) {
							log := ""
							// 构建Pod只有一个容器，所以取第一个容器的日志即构建日志
							logBytes, err := client.GetPodLogs(ctx, node.K8sWorkerSpace, pod.Name, pod.Containers[0].Name, nil, false)
							if err == nil {
								log = string(logBytes)
							}

							if err := repositories.UpdatePipeline(pod.BuildId, rlt, &log); err == nil && pipeline.IsComplete(rlt) {
								// 清理Pod
								client.DeletePod(context.TODO(), node.K8sWorkerSpace, pod.Name)
								delete(allK8sPipelines, pod.BuildId)

								// 通知制品监视器流水线构建完成
								artifactWatcher <- pod.BuildId
							}
						}
					}
				}
			}
		}
	}

	// 查找当前构建机节点中没有找到的构建任务，标记为Interrupt
	// 1. 查找所有未完成的流水线
	if pipelines, err := repositories.QueryIncompletePipeline(); err == nil {
		// 2. 如果这些流水线在pod中存在，则忽略，如果不存在，则标记为interrupt
		for _, p := range pipelines {
			if _, ok := currentExistsPipeline[p.LastRunId]; !ok {
				if err := repositories.UpdatePipeline(p.LastRunId, pipeline.BuildingInterrupt, nil); err != nil {
					fmt.Println(err.Error())
				}
			}
		}
	}
}
