package kube

import (
	"golang.org/x/net/context"
	av1 "k8s.io/api/apps/v1"
	cv1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	"strings"
)

// ApplyDeployment kubectl apply -f yml
func (client *Client) ApplyDeployment(namespace *string, yml *appsv1.DeploymentApplyConfiguration) (*av1.Deployment, error) {
	if _, err := client.getOrCreateNamespace(namespace); err != nil {
		return nil, err
	}
	return client.clientSet.AppsV1().Deployments(*namespace).Apply(context.TODO(), yml, *client.defaultApplyOptions)
}

// ApplyService kubectl apply -f yml
func (client *Client) ApplyService(namespace *string, yml *corev1.ServiceApplyConfiguration) (*cv1.Service, error) {
	if _, err := client.getOrCreateNamespace(namespace); err != nil {
		return nil, err
	}
	return client.clientSet.CoreV1().Services(*namespace).Apply(context.TODO(), yml, *client.defaultApplyOptions)
}

// ApplyPod kubectl apply -f yml
func (client *Client) ApplyPod(namespace *string, yml *corev1.PodApplyConfiguration) (*cv1.Pod, error) {
	if _, err := client.getOrCreateNamespace(namespace); err != nil {
		return nil, err
	}
	return client.clientSet.CoreV1().Pods(*namespace).Apply(context.TODO(), yml, *client.defaultApplyOptions)
}

const namespace_yml = `
apiVersion: v1
kind: Namespace
metadata:
  name: {{.Namespace}}
`

// getOrCreateNamespace 获取或创建namespace
func (client *Client) getOrCreateNamespace(namespace *string) (*cv1.Namespace, error) {

	cfg := strings.ReplaceAll(namespace_yml, "{{.Namespace}}", *namespace)
	yml := corev1.NamespaceApplyConfiguration{}
	err := DecodeYaml(&cfg, &yml)

	if err != nil {
		return nil, err
	}

	return client.clientSet.CoreV1().Namespaces().Apply(context.TODO(), &yml, *client.defaultApplyOptions)
}
