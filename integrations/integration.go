package integrations

import "os/exec"

type Integration interface {
	PrepareCommand(taskName string) *exec.Cmd
	GetTaskList() []string
}