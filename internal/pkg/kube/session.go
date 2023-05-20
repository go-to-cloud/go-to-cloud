package kube

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"k8s.io/client-go/tools/remotecommand"
)

const (
	endOfTransmission = "\u0004"
)

type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
	Done() chan struct{}
}

// TerminalSession implements PtyHandler
type TerminalSession struct {
	WsConn   *websocket.Conn
	SizeChan chan remotecommand.TerminalSize
	DoneChan chan struct{}
}

func (t *TerminalSession) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.SizeChan:
		return &size
	case <-t.DoneChan:
		return nil
	}
}

func (t *TerminalSession) Done() chan struct{} {
	return t.DoneChan
}

type TerminalMessage struct {
	Operation string `json:"operation"`
	Data      string `json:"data"`
	Rows      uint16 `json:"rows"`
	Cols      uint16 `json:"cols"`
}

func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.WsConn.ReadMessage()
	if err != nil {
		return copy(p, endOfTransmission), err
	}
	var msg TerminalMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		return copy(p, endOfTransmission), err
	}
	switch msg.Operation {
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		t.SizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	default:
		return copy(p, endOfTransmission), fmt.Errorf("unknown message type '%s'", msg.Operation)
	}
}

func (t *TerminalSession) Write(p []byte) (int, error) {
	msg, err := json.Marshal(TerminalMessage{
		Operation: "stdout",
		Data:      string(p),
	})
	if err != nil {
		return 0, err
	}
	if err := t.WsConn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return 0, err
	}
	return len(p), nil
}

func (t *TerminalSession) Close() error {
	return t.WsConn.Close()
}
