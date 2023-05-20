package builder

import (
	"context"
	"errors"
	"github.com/patrickmn/go-cache"
	"go-to-cloud/internal/models/builder"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"strconv"
	"time"
)

var idleNodes *cache.Cache

// GetWorkingNodes 获取正在工作的节点数量
func GetWorkingNodes(workerId uint) (int, error) {
	node, err := repositories.GetBuildNodesById(workerId)
	if err != nil {
		return 0, err
	}

	if node == nil {
		return 0, errors.New("没有找到构建节点配置")
	}

	if node.NodeType == int(builder.K8s) {
		return func() (int, error) {
			if a, e := tryGetPodStatusFromCache(node, getPodDescription); e != nil {
				return 0, e
			} else {
				return len(a), nil
			}
		}()
	}

	return 0, errors.New("不支持的构建节点类型")
}

func getPodDescription(node *repositories.BuilderNode) ([]kube.PodDescription, error) {
	client, err := kube.NewClient(node.DecryptKubeConfig())
	if err != nil {
		return nil, err
	}

	rlt, err := client.GetPods(context.TODO(), node.K8sWorkerSpace, BuildIdSelectorLabel, func() string {
		return "builder=" + NodeSelectorLabel
	}, true)

	if err != nil {
		return nil, err
	} else {
		return kube.TrimPodDetailDescriptions(rlt), nil
	}
}

func tryGetPodStatusFromCache(node *repositories.BuilderNode, f func(node *repositories.BuilderNode) ([]kube.PodDescription, error)) ([]kube.PodDescription, error) {
	if v, ok := idleNodes.Get(strconv.Itoa(int(node.ID))); ok {
		return v.([]kube.PodDescription), nil
	} else {
		if n, e := f(node); e != nil {
			return nil, e
		} else {
			idleNodes.Set(strconv.Itoa(int(node.ID)), n, cache.DefaultExpiration)
			return n, nil
		}
	}
}

func init() {
	idleNodes = cache.New(time.Minute, 0)
}
