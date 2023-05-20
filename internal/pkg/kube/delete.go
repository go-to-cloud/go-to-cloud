package kube

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Delete 删除
func (client *Client) Delete(ns, deployment *string) error {
	return client.clientSet.AppsV1().Deployments(*ns).Delete(context.TODO(), *deployment, metav1.DeleteOptions{})
}
