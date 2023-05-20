package kube

import (
	"fmt"
	"golang.org/x/net/context"
	cv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"strings"
)

func (client *Client) GetAllNamespaces(includeKube bool) ([]string, error) {

	if all, err := client.clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{}); err != nil {
		return nil, err
	} else {
		ns := make([]string, 0)
		for _, namespace := range all.Items {
			if !includeKube {
				if strings.HasPrefix(namespace.Name, "kube-") || strings.HasPrefix(namespace.Name, "kubernetes-") {
					continue
				}
			}
			ns = append(ns, namespace.Name)
		}
		return ns, nil
	}
}

// GetOrAddNamespace 获取或创建名字空间
func (client *Client) GetOrAddNamespace(ns *string) (*cv1.Namespace, error) {

	kind := "Namespace"
	apiVer := "v1"
	namespace := corev1.NamespaceApplyConfiguration{
		TypeMetaApplyConfiguration: v1.TypeMetaApplyConfiguration{
			Kind:       &kind,
			APIVersion: &apiVer,
		},
		ObjectMetaApplyConfiguration: &v1.ObjectMetaApplyConfiguration{
			Name: ns,
		},
	}
	rlt, err := client.clientSet.CoreV1().Namespaces().Apply(context.TODO(), &namespace, *client.defaultApplyOptions)

	if err != nil {
		fmt.Println(err)
	}
	return rlt, err
}

func (client *Client) DeleteNamespace(ns *string) error {
	deletePolicy := metav1.DeletePropagationForeground
	return client.clientSet.CoreV1().Namespaces().Delete(context.TODO(), *ns, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}
