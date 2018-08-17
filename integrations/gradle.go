/*
Sniperkit-Bot
- Status: analyzed
*/

package integrations

import (
	"bufio"
	"bytes"
	"os/exec"
	"regexp"

	"github.com/sniperkit/snk.fork.webtaskrunner/config"
)

//NewGradleIntegration returns a new instance of the Gradle integration wrapper.
func NewGradleIntegration(config *config.GradleConfig) *GradleIntegration {
	return &GradleIntegration{
		config: config,
	}
}

//GradleIntegration implements the integration interface.
type GradleIntegration struct {
	config *config.GradleConfig
}

func (i *GradleIntegration) getGradleFileParameters() (string, string) {
	if i.config.ExecutionDir != "" {
		return "--project-dir", i.config.ExecutionDir
	}
	return "", ""
}

//PrepareCommand prepares an exec.Cmd so that it will start the given task when executed
func (i *GradleIntegration) PrepareCommand(taskName string) *exec.Cmd {
	gradleProjectDirFlag, gradleProjectDir := i.getGradleFileParameters()
	cmd := exec.Command("gradle", gradleProjectDirFlag, gradleProjectDir, taskName)
	return cmd
}

//GetTaskList returns as list of tasks
func (i *GradleIntegration) GetTaskList() []TaskInfo {
	gradleProjectDirFlag, gradleProjectDir := i.getGradleFileParameters()
	stdOutBytes, err := exec.Command("gradle", gradleProjectDirFlag, gradleProjectDir, "tasks").Output()

	if err != nil {
		panic(err)
	}

	taskInfoList := []TaskInfo{}
	scanner := bufio.NewScanner(bytes.NewBuffer(stdOutBytes))
	for scanner.Scan() {
		stdOutLine := scanner.Text()

		taskInfo := i.extractTaskInfo(stdOutLine)
		if taskInfo != nil {
			taskInfo.IntegrationName = "gradle"
			taskInfoList = append(taskInfoList, *taskInfo)
		}
	}

	return taskInfoList
}

func (i *GradleIntegration) extractTaskInfo(s string) *TaskInfo {

	regexps := []*regexp.Regexp{
		regexp.MustCompile("^((?P<taskName>\\w+)( - (?P<description>.*)))$"),
		regexp.MustCompile("^(?P<taskName>\\w+)$"),
	}

	for _, re := range regexps {
		if re.MatchString(s) {
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
	}

	return nil
}
