package taskrunnerlisthandler

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/config"
	"net/http"
)

// New creates a new instance of TaskRunnersListHandler
func New(frontendConfiguration []*config.FrontendInfo) *TaskRunnersListHandler {
	return &TaskRunnersListHandler{
		frontendConfiguration: frontendConfiguration,
	}
}

// TaskRunnersListHandler provides a simple ajax json api which queries all taskrunners returns them as json array
type TaskRunnersListHandler struct {
	frontendConfiguration []*config.FrontendInfo
}

// ServeHTTP implements the http.Handler interface
func (h *TaskRunnersListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pJson, err := json.Marshal(h.frontendConfiguration)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(pJson)
}
