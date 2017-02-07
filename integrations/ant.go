package integrations

import (
	"os/exec"
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func NewAntIntegration() *AntIntegration {
	return &AntIntegration{}
}

type AntIntegration struct{}

func (i *AntIntegration) PrepareCommand(taskName string) *exec.Cmd {
	cmd := exec.Command("ant", taskName)
	return cmd
}

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
