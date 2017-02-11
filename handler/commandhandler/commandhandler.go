package commandhandler

import (
	"bytes"
	"encoding/json"
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

	outChannel := make(chan byte)
	errChannel := make(chan error)
	finished := false
	go execution.ExecuteTask(taskName, h.integration, outChannel, errChannel)

	bb := bytes.NewBuffer([]byte{})
	for {
		select {
		case outByte, ok := <-outChannel:
			if ok {
				if ok {
					bb.WriteByte(outByte)
					if outByte == '\n' || outByte == '\r' {
						h.writeLogLine(bb.Bytes(), ws)
						bb.Reset()
					}
				}
			} else {
				h.writeLogLine(bb.Bytes(), ws)
				bb.Reset()
				finished = true
				break
			}

		case err, ok := <-errChannel:
			if ok {
				h.writeErrorLine(err, ws)
				break
			}
		}

		if finished {
			break
		}
	}
}

func (h *WebSocketHandler) writeErrorLine(err error, ws *websocket.Conn) {
	responseLine := responseError{
		Status: 2,
		Error:  string(err.Error()),
	}

	jsonResponseLine, _ := json.Marshal(responseLine)
	ws.Write(jsonResponseLine)
}
func (h *WebSocketHandler) writeLogLine(lineBytes []byte, ws *websocket.Conn) {
	responseLine := responseLine{
		Status: 1,
		Line:   string(lineBytes),
	}

	jsonResponseLine, _ := json.Marshal(responseLine)
	ws.Write(jsonResponseLine)
}
