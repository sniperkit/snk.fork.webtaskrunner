/*
Sniperkit-Bot
- Status: analyzed
*/

package integrations

import (
	"bufio"
	"bytes"
	"os/exec"
	"path"
	"strings"

	"github.com/sniperkit/snk.fork.webtaskrunner/config"
)

//NewGulpIntegration returns a new instance of the gulp integration wrapper.
func NewGulpIntegration(config *config.GulpConfig) *GulpIntegration {
	return &GulpIntegration{
		config: config,
	}
}

//GulpIntegration implements the integration interface.
type GulpIntegration struct {
	config *config.GulpConfig
}

func (i *GulpIntegration) getGulpFileParameters() (string, string) {
	if i.config.GulpFilePath != "" {
		return "--gulpfile", path.Join(i.config.GulpFilePath, "gulpfile.js")
	}
	return "", ""
}

//PrepareCommand prepares an exec.Cmd so that it will start the given task when executed
func (i *GulpIntegration) PrepareCommand(taskName string) *exec.Cmd {
	gulpFileFlag, gulpFilePath := i.getGulpFileParameters()
	cmd := exec.Command("gulp", gulpFileFlag, gulpFilePath, taskName)
	cmd.Dir = i.config.ExecutionDir
	return cmd
}

//GetTaskList returns as list of tasks
func (i *GulpIntegration) GetTaskList() []TaskInfo {
	cmd := exec.Command("gulp", "--tasks-simple")
	cmd.Dir = i.config.ExecutionDir
	stdOutBytes, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	taskInfoList := []TaskInfo{}
	scanner := bufio.NewScanner(bytes.NewBuffer(stdOutBytes))
	for scanner.Scan() {
		stdOutLine := scanner.Text()
		taskInfo := TaskInfo{
			TaskName:        strings.Trim(stdOutLine, " "),
			Description:     "",
			IntegrationName: "gulp",
		}
		taskInfoList = append(taskInfoList, taskInfo)
	}

	return taskInfoList
}
