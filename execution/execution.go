package execution

import (
	"bytes"
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"os/exec"
)

// ExecuteTask executes the given taskName using the given integration
// the stdout of the task will be sent continuously to stdoutChannel
func ExecuteTask(taskName string, integration integrations.Integration, stdoutChannel chan string) {

	cmd := integration.PrepareCommand(taskName)

	stdErrBytes := []byte{}
	stdErrWriter := bytes.NewBuffer(stdErrBytes)
	stdOutBytes := []byte{}
	stdOutWriter := bytes.NewBuffer(stdOutBytes)

	cmd.Stderr = stdErrWriter
	cmd.Stdout = stdOutWriter
	errChannel := make(chan error)

	go runCommand(cmd, errChannel)

	for {
		select {
		case err := <-errChannel:
			if err != nil {
				fmt.Println("ERROR !!!", err)
			} else {
				fmt.Println("Finished")
			}
			processOutputs(stdOutWriter, stdErrWriter, stdoutChannel)
			close(stdoutChannel)
			return
		default:
			processOutputs(stdOutWriter, stdErrWriter, stdoutChannel)
		}
	}

}

func runCommand(cmd *exec.Cmd, errChannel chan error) {
	err := cmd.Run()
	if err != nil {
		errChannel <- err
	}

	close(errChannel)
}

func processOutputs(stdOutWriter *bytes.Buffer, stdErrWriter *bytes.Buffer, stdoutChannel chan string) {
	if stdOutWriter.Len() > 0 {
		b := make([]byte, stdOutWriter.Len())
		_, err := stdOutWriter.Read(b)
		if err != nil {
			panic(err)
		}
		stdoutChannel <- string(b)
	}
	if stdErrWriter.Len() > 0 {
		b := make([]byte, stdErrWriter.Len())
		_, err := stdErrWriter.Read(b)
		if err != nil {
			panic(err)
		}
		stdoutChannel <- string(b)
	}
}
