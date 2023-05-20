package k8s

import (
	"fmt"
	"go-to-cloud/internal/models/deploy/k8s"
	"go-to-cloud/internal/pkg/kube"
)

// Ping 测试K8s服务是否可用
func Ping(testing *k8s.Testing) (string, error) {

	c, e := kube.NewClient(testing.KubeConfig)
	if e != nil {
		return "", e
	}

	info, e := c.GetClientSet().Discovery().ServerVersion()

	return fmt.Sprintf("%s.%s", info.Major, info.Minor), e
}
