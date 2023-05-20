package kube

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

const RestartAt = "restart-at"

// Restart 重启
func (client *Client) Restart(ns, deployment *string) error {

	deploy, err := client.clientSet.AppsV1().Deployments(*ns).Get(context.TODO(), *deployment, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if deploy.Spec.Template.ObjectMeta.Annotations == nil {
		deploy.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	deploy.Spec.Template.ObjectMeta.Annotations[RestartAt] = time.Now().String()

	_, err = client.clientSet.AppsV1().Deployments(*ns).Update(context.TODO(), deploy, metav1.UpdateOptions{})

	return err
}
