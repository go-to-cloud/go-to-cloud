package uninstall

import (
	"errors"
	"go-to-cloud/internal/repositories"
)

// OnK8s 卸载K8s节点
func OnK8s(userId, nodeId uint) error {
	if userId <= 0 || nodeId <= 0 {
		return errors.New("not allowed")
	}
	return repositories.DeleteBuilderNode(userId, nodeId)
}
