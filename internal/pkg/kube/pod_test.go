package kube

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetPod(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	restcfg, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	assert.NoError(t, err)

	client, err := NewClientByRestConfig(restcfg)
	assert.NoError(t, err)

	pods, err := client.GetPods(context.TODO(), "kube-system", "", func() string { return "" }, true)
	assert.NoError(t, err)
	assert.NotEmpty(t, pods)

	pod := func() *PodDescription {
		for i, p := range pods {
			if p.Name == "kube-apiserver-docker-desktop" {
				return &pods[i].PodDescription
			}
		}
		return nil
	}()
	containerName := "kube-apiserver"
	var tailLine int64 = 1024
	log, err := client.GetPodStreamLogs(context.TODO(), "kube-system", pod.Name, containerName, &tailLine, true, false)
	assert.NoError(t, err)
	defer log.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, log)

	logBuilder := strings.Builder{}
	logBuilder.WriteString(buf.String())

	content := make([]byte, 1024)
	for {
		n, err := log.Read(content)
		assert.NoError(t, err)
		msg := string(content[:n])
		assert.NotNil(t, msg)
		logBuilder.WriteString(msg + "\n")
	}
}
