package taskshandler

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"golang.org/x/net/websocket"
	"net/http"
)

// New returns a new instance of TaskListWebSocketHandler
func New(integrationList []integrations.Integration) *TaskListWebSocketHandler {
	return &TaskListWebSocketHandler{
		integrationList: integrationList,
	}
}

// TaskListWebSocketHandler provides an websocket connection over which a task command may be started on the server.
type TaskListWebSocketHandler struct {
	integrationList []integrations.Integration
}

// GetHandler returns the websocket http.Handler
func (h *TaskListWebSocketHandler) GetHandler() http.Handler {
	return websocket.Handler(h.commandHandler)
}

func (h *TaskListWebSocketHandler) commandHandler(ws *websocket.Conn) {
	fmt.Printf("query all tasks from all integrations\n")

	finishedChannel := make(chan bool)
	taskInfoChannel := make(chan integrations.TaskInfo)

	for _, integration := range h.integrationList {
		integration.GetTaskList()
		go func(integration integrations.Integration, taskInfoChannel chan integrations.TaskInfo, finishedChannel chan bool) {
			taskInfoList := integration.GetTaskList()
			for _, taskInfo := range taskInfoList {
				taskInfoChannel <- taskInfo
			}
			finishedChannel <- true
		}(integration, taskInfoChannel, finishedChannel)
	}

	go func(taskInfoChannel chan integrations.TaskInfo) {
		for {
			select {
			case taskInfo, ok := <-taskInfoChannel:
				if ok {
					h.writeTaskInfo(taskInfo, ws)
				} else {
					fmt.Println("TaskInfoChannelClosed")
					return
				}
			}
		}

	}(taskInfoChannel)

	for i := 0; i < len(h.integrationList); i++ {
		<-finishedChannel
		fmt.Println("finished query")
	}
	close(taskInfoChannel)
	fmt.Println("finished all task query")

}

func (h *TaskListWebSocketHandler) writeTaskInfo(taskInfo integrations.TaskInfo, ws *websocket.Conn) {

	jsonResponseLine, _ := json.Marshal(taskInfo)
	ws.Write(jsonResponseLine)
}
