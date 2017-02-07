package task_list_handler

import (
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"fmt"
	"encoding/json"
	"net/http"
)

func New(integration integrations.Integration) *TaskListHandler {
	return &TaskListHandler{
		integration : integration,
	}
}

type TaskListHandler struct {
	integration integrations.Integration
}

func (h *TaskListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tasks := h.integration.GetTaskList()
	pJson, err := json.Marshal(tasks)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(pJson)
}

