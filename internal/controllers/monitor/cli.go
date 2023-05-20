package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-to-cloud/internal/pkg/response"
	"go-to-cloud/internal/services/monitor"
	"net/http"
)

func getWebSocketParams(ctx *gin.Context) (ws *websocket.Conn, k8sRepoId, deploymentId uint, podName, containerName string, err error) {

	k8sRepoId, err = getUIntParamFromQueryOrPath("k8s", ctx, false)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	deploymentId, err = getUIntParamFromQueryOrPath("deploymentId", ctx, false)
	if err != nil {
		response.BadRequest(ctx, err.Error())
		return
	}
	podName = ctx.Param("podName")         // 允许为空，空时进入默认第一个容器内部
	containerName = ctx.Query("container") // 允许为空，空时进入默认第一个容器内部

	ws, err = monitor.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	return
}

// DisplayLog 查看容器日志
func DisplayLog(ctx *gin.Context) {
	ws, k8sRepoId, deploymentId, podName, containerName, err := getWebSocketParams(ctx)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	defer func() {
		ws.Close()
	}()

	monitor.XTermInteractiveLogs(ws, k8sRepoId, deploymentId, podName, containerName, ctx.Done())
}

// Interactive 进入容器内部执行命令行交互
func Interactive(ctx *gin.Context) {
	ws, k8sRepoId, deploymentId, podName, containerName, err := getWebSocketParams(ctx)

	if err != nil {
		msg := err.Error()
		response.Fail(ctx, http.StatusInternalServerError, &msg)
		return
	}

	monitor.XTermInteractiveShell(ws, k8sRepoId, deploymentId, podName, containerName, ctx.Done())
}
