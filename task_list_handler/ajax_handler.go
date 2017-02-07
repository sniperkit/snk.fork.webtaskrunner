package task_list_handler

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"net/http"
)

// New creates a new instance of TaskListHandler
func New(integration integrations.Integration) *TaskListHandler {
	return &TaskListHandler{
		integration: integration,
	}
}

// TaskListHandler provides a simple ajax json api which queries all tasks of an integration and returns them
// as json array
type TaskListHandler struct {
	integration integrations.Integration
}

// ServeHTTP implements the http.Handler interface
func (h *TaskListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tasks := h.integration.GetTaskList()
	pJson, err := json.Marshal(tasks)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(pJson)
}
