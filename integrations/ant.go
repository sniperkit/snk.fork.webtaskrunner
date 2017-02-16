package integrations

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/config"
	"os/exec"
	"strings"
)

//NewAntIntegration returns a new instance of the ant integration wrapper.
func NewAntIntegration(config *config.AntConfig) *AntIntegration {
	return &AntIntegration{
		config: config,
	}
}

//AntIntegration implements the integration interface.
type AntIntegration struct {
	config *config.AntConfig
}

//PrepareCommand prepares an exec.Cmd so that it will start the given task when executed
func (i *AntIntegration) PrepareCommand(taskName string) *exec.Cmd {
	cmd := exec.Command("ant", taskName)
	return cmd
}

//GetTaskList returns as list of tasks
func (i *AntIntegration) GetTaskList() []TaskInfo {
	stdOutBytes, err := exec.Command("ant", "-p", "build.xml").Output()

	if err != nil {
		panic(err)
	}

	taskInfoList := []TaskInfo{}
	scanner := bufio.NewScanner(bytes.NewBuffer(stdOutBytes))
	for scanner.Scan() {
		stdOutLine := scanner.Text()
		fmt.Println(stdOutLine)
		if strings.Contains(stdOutLine, "Buildfile:") {
			continue
		}
		if strings.Contains(stdOutLine, "Main targets:") {
			continue
		}
		if strings.Contains(stdOutLine, "Other targets:") {
			continue
		}
		if len(stdOutLine) == 0 {
			continue
		}

		taskInfo := TaskInfo{
			TaskName:        strings.Trim(stdOutLine, " "),
			Description:     "",
			IntegrationName: "ant",
		}
		taskInfoList = append(taskInfoList, taskInfo)
	}

	return taskInfoList
}
