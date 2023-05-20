package monitor

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"go-to-cloud/internal/pkg/kube"
	"go-to-cloud/internal/repositories"
	"k8s.io/client-go/tools/remotecommand"
)

func newTerminalSession(ws *websocket.Conn) (*kube.TerminalSession, error) {
	session := &kube.TerminalSession{
		WsConn:   ws,
		SizeChan: make(chan remotecommand.TerminalSize),
		DoneChan: make(chan struct{}),
	}
	return session, nil
}

func XTermInteractiveShell(ws *websocket.Conn, k8sRepoId, deploymentId uint, podName, containerName string, cancel <-chan struct{}) {

	c := context.Background()

	go func() {
		<-cancel
		c.Done()
	}()

	session, err := newTerminalSession(ws)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer session.Close()

	repo, err := repositories.QueryK8sRepoById(k8sRepoId)
	if err != nil {
		fmt.Println(err)
		return
	}

	deployment, err := repositories.GetDeploymentById(deploymentId)
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := kube.NewClient(&repo.KubeConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := client.Shell(c, deployment.K8sNamespace, podName, containerName, session); err != nil {
		session.Done()
	}
}
