package monitor

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func XTermInteractiveLogs(ws *websocket.Conn, k8sRepoId, deploymentId uint, podName, containerName string, cancel <-chan struct{}) {
	c := context.Background()

	go func() {
		<-cancel
		c.Done()
	}()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read error")
			break
		}
		switch mt {
		case websocket.TextMessage:
			containerName = string(message)
			c.Done()
			go followLogs(c, k8sRepoId, deploymentId, podName, containerName, false, func(log []byte) {
				ws.WriteMessage(websocket.TextMessage, log)
			})

		case websocket.PingMessage:
			_ = k8sRepoId
			_ = containerName
			err = ws.WriteMessage(websocket.PongMessage, []byte("pong"))
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
