package integrations

import (
	"os/exec"
	"bufio"
	"bytes"
	"regexp"
)

//NewGradleIntegration returns a new instance of the Gradle integration wrapper.
func NewGradleIntegration() *GradleIntegration {
	return &GradleIntegration{}
}

//GradleIntegration implements the integration interface.
type GradleIntegration struct{}

//PrepareCommand prepares an exec.Cmd so that it will start the given task when executed
func (i *GradleIntegration) PrepareCommand(taskName string) *exec.Cmd {
	cmd := exec.Command("gradle", taskName)
	return cmd
}

//GetTaskList returns as list of tasks
func (i *GradleIntegration) GetTaskList() []string {
	stdOutBytes, err := exec.Command("gradle", "tasks").Output()

	if err != nil {
		panic(err)
	}

	targets := []string{}
	scanner := bufio.NewScanner(bytes.NewBuffer(stdOutBytes))
	for scanner.Scan() {
		stdOutLine := scanner.Text()

		taskInfo := i.extractTaskInfo(stdOutLine)
		if taskInfo != nil {
			targets = append(targets, taskInfo.TaskName)
		}
	}

	return targets
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