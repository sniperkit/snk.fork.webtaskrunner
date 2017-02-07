package websocket_handler

import (
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"golang.org/x/net/websocket"
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/execution"
	"net/http"
	"log"
)

func New(integration integrations.Integration) *WebSocketHandler {
	return &WebSocketHandler{
		integration : integration,
	}
}

type WebSocketHandler struct {
	integration integrations.Integration
}

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

