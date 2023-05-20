package kube

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	clientSet           *kubernetes.Clientset
	defaultApplyOptions *meta.ApplyOptions
	config              *rest.Config
}

func (client *Client) GetClientSet() *kubernetes.Clientset {
	return client.clientSet
}

func newClient(c *kubernetes.Clientset, config *rest.Config) (*Client, error) {
	m := meta.ApplyOptions{
		FieldManager: "application/apply-patch+yaml",
		Force:        true,
	}

	client := Client{
		clientSet:           c,
		defaultApplyOptions: &m,
		config:              config,
	}
	return &client, nil

}

func NewClientByRestConfig(cfg *rest.Config) (*Client, error) {
	if c, e := kubernetes.NewForConfig(cfg); e != nil {
		return nil, e
	} else {
		return newClient(c, cfg)
	}
}

func NewClientFromToken(token, host *string) (*Client, error) {
	return NewClientByRestConfig(&rest.Config{
		BearerToken:     *token,
		TLSClientConfig: rest.TLSClientConfig{Insecure: true},
		Host:            *host,
	})
}

// NewClient 创建k8s客户端对象
func NewClient(config *string) (*Client, error) {
	kubeConfig, err := clientcmd.BuildConfigFromKubeconfigGetter("", func() (*api.Config, error) {
		return clientcmd.Load([]byte(*config))
	})
	if err != nil {
		return nil, err
	}

	if c, e := kubernetes.NewForConfig(kubeConfig); e != nil {
		return nil, e
	} else {
		return newClient(c, kubeConfig)
	}
}
