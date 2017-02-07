package integrations

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

//NewAntIntegration returns a new instance of the ant integration wrapper.
func NewAntIntegration() *AntIntegration {
	return &AntIntegration{}
}

//AntIntegration implements the integration interface.
type AntIntegration struct{}

//PrepareCommand prepares an exec.Cmd so that it will start the given task when executed
func (i *AntIntegration) PrepareCommand(taskName string) *exec.Cmd {
	cmd := exec.Command("ant", taskName)
	return cmd
}

//GetTaskList returns as list of tasks
func (i *AntIntegration) GetTaskList() []string {
	stdOutBytes, err := exec.Command("ant", "-p", "build.xml").Output()

	if err != nil {
		panic(err)
	}

	targets := []string{}
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

		targets = append(targets, strings.Trim(stdOutLine, " "))
	}

	return targets
}
