package integrations

import "os/exec"

//Integration defines the interface of a build system integration
type Integration interface {
	PrepareCommand(taskName string) *exec.Cmd
	GetTaskList() []string
}
