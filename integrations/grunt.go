/*
Sniperkit-Bot
- Status: analyzed
*/

package integrations

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/sniperkit/snk.fork.webtaskrunner/config"
)

//NewGruntIntegration returns a new instance of the grunt integration wrapper.
func NewGruntIntegration(config *config.GruntConfig) *GruntIntegration {
	return &GruntIntegration{
		config: config,
	}
}

//GruntIntegration implements the integration interface.
type GruntIntegration struct {
	config *config.GruntConfig
}

func (i *GruntIntegration) getGruntFileParameters() (string, string) {
	if i.config.GruntFilePath != "" {
		return "--gruntfile", path.Join(i.config.GruntFilePath, "Gruntfile.js")
	}
	return "", ""
}

//PrepareCommand prepares an exec.Cmd so that it will start the given task when executed
func (i *GruntIntegration) PrepareCommand(taskName string) *exec.Cmd {
	gruntFileFlag, gruntFilePath := i.getGruntFileParameters()
	cmd := exec.Command("grunt", gruntFileFlag, gruntFilePath, taskName)
	cmd.Dir = i.config.ExecutionDir
	return cmd
}

//GetTaskList returns as list of tasks
func (i *GruntIntegration) GetTaskList() []TaskInfo {
	gruntFileFlag, gruntFilePath := i.getGruntFileParameters()
	cmd := exec.Command("grunt", gruntFileFlag, gruntFilePath, "--help")
	cmd.Dir = i.config.ExecutionDir
	stdOutBytes, err := cmd.Output()
	if err != nil {
		fmt.Println(string(stdOutBytes))
		return []TaskInfo{}
	}
	taskInfoList := i.scanTasksFromOutput(stdOutBytes)

	return taskInfoList
}

func (i *GruntIntegration) scanTasksFromOutput(stdOutBytes []byte) []TaskInfo {

	bScanTasks := false
	taskInfoList := []TaskInfo{}
	scanner := bufio.NewScanner(bytes.NewBuffer(stdOutBytes))
	for scanner.Scan() {
		stdOutLine := scanner.Text()
		fmt.Println(stdOutLine)
		if strings.Contains(stdOutLine, "Available tasks") {
			bScanTasks = true
			continue
		}
		if bScanTasks {
			if stdOutLine == "" {
				bScanTasks = false
			} else {
				taskInfo := i.parseTaskInfo(stdOutLine)
				if taskInfo != nil {
					taskInfo.IntegrationName = "grunt"
					taskInfoList = append(taskInfoList, *taskInfo)
				}
			}
		}
	}

	return taskInfoList
}

func (i *GruntIntegration) parseTaskInfo(s string) *TaskInfo {
	re := regexp.MustCompile("^\\s*(?P<taskName>\\w+)\\s*(?P<description>.*)$")
	if !re.MatchString(s) {
		return nil
	}
	subMatches := re.FindAllStringSubmatch(s, -1)[0]
	names := re.SubexpNames()
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
