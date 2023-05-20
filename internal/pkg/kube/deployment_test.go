package kube

import (
	"context"
	"github.com/stretchr/testify/assert"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"testing"
)

func TestListDeployments(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	restcfg, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	assert.NoError(t, err)
	kubeClient, err := kubernetes.NewForConfig(restcfg)
	assert.NoError(t, err)

	client := &Client{
		clientSet: kubeClient,
		defaultApplyOptions: &meta.ApplyOptions{
			FieldManager: "application/apply-patch+yaml",
			Force:        true,
		},
	}

	m := make(map[uint]bool)
	m[1] = true
	lst, err2 := client.GetDeployments(context.TODO(), 1, "default", &m, false)
	assert.NoError(t, err2)

	assert.NotEmpty(t, lst)
}
