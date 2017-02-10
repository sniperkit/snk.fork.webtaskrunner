package integrations

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

//NewGruntIntegration returns a new instance of the grunt integration wrapper.
func NewGruntIntegration() *GruntIntegration {
	return &GruntIntegration{}
}

//GruntIntegration implements the integration interface.
type GruntIntegration struct{}

//PrepareCommand prepares an exec.Cmd so that it will start the given task when executed
func (i *GruntIntegration) PrepareCommand(taskName string) *exec.Cmd {
	cmd := exec.Command("grunt", taskName)
	return cmd
}

//GetTaskList returns as list of tasks
func (i *GruntIntegration) GetTaskList() []string {
	stdOutBytes, err := exec.Command("grunt", "--help").Output()

	if err != nil {
		panic(err)
	}

	bScanTasks := false
	targets := []string{}
	scanner := bufio.NewScanner(bytes.NewBuffer(stdOutBytes))
	for scanner.Scan() {
		stdOutLine := scanner.Text()
		//fmt.Println(stdOutLine)
		if strings.Contains(stdOutLine, "Available tasks") {
			bScanTasks = true
			continue
		}
		if bScanTasks {
			if stdOutLine == "" {
				bScanTasks = false
			} else {
				taskInfo := i.parseTaskInfo(stdOutLine)
				targets = append(targets, taskInfo.TaskName)
			}
		}
	}

	return targets
}

func (i *GruntIntegration) parseTaskInfo(s string) *TaskInfo {
	re := regexp.MustCompile("^\\s*(?P<taskName>\\w+)\\s*(?P<description>.*)$")
	subMatches := re.FindAllStringSubmatch(s, -1)[0]
	fmt.Println(subMatches)
	names := re.SubexpNames()
	fmt.Println(names)
	taskInfo := TaskInfo{}
	for i, n := range subMatches {
		name := names[i]
		if name == "" {
			continue
		}

		if name == "taskName" {
			taskInfo.TaskName = n
		}
		if name == "description" {
			taskInfo.Description = n
		}
	}
	return &taskInfo
}
