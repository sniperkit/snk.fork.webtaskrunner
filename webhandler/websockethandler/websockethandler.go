package websockethandler

import (
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/execution"
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

// New returns a new instance of WebSocketHandler
func New(integration integrations.Integration) *WebSocketHandler {
	return &WebSocketHandler{
		integration: integration,
	}
}

// WebSocketHandler provides an websocket connection over which a task command may be started on the server.
type WebSocketHandler struct {
	integration integrations.Integration
}

// GetHandler returns the websocket http.Handler
func (h *WebSocketHandler) GetHandler() http.Handler {
	return websocket.Handler(h.commandHandler)
}

func (h *WebSocketHandler) commandHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	taskName := string(msg[:n])
	fmt.Printf("start task: %s\n", taskName)

	outChannel := make(chan string)
	go execution.ExecuteTask(taskName, h.integration, outChannel)

	for {
		if out, ok := <-outChannel; ok {
			if ok {
				ws.Write([]byte(out))
			}
		} else {
			break
		}
	}
}
