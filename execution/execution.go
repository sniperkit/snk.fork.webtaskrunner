/*
Sniperkit-Bot
- Status: analyzed
*/

package execution

import (
	"bytes"
	"os/exec"

	"github.com/sniperkit/snk.fork.webtaskrunner/integrations"
)

// ExecuteTask executes the given taskName using the given integration
// the stdout of the task will be sent continuously to stdoutChannel
func ExecuteTask(taskName string, integration integrations.Integration, stdoutChannel chan byte, errorChannel chan error) {

	cmd := integration.PrepareCommand(taskName)

	stdErrBytes := []byte{}
	stdErrWriter := bytes.NewBuffer(stdErrBytes)
	stdOutBytes := []byte{}
	stdOutWriter := bytes.NewBuffer(stdOutBytes)

	cmd.Stderr = stdErrWriter
	cmd.Stdout = stdOutWriter
	errChannel := make(chan error)
	var errValue error
	finished := false
	go runCommand(cmd, errChannel)

	for {
		select {
		case err := <-errChannel:
			if err != nil {
				errValue = err
			}
			processOutputs(stdOutWriter, stdErrWriter, stdoutChannel)
			finished = true
			break
		default:
			processOutputs(stdOutWriter, stdErrWriter, stdoutChannel)
		}

		if finished {
			if errValue != nil {
				errorChannel <- errValue
			}
			close(errorChannel)
			close(stdoutChannel)

			break
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

func processOutputs(stdOutWriter *bytes.Buffer, stdErrWriter *bytes.Buffer, stdoutChannel chan byte) {
	if stdOutWriter.Len() > 0 {
		bytes := make([]byte, stdOutWriter.Len())
		_, err := stdOutWriter.Read(bytes)
		if err != nil {
			panic(err)
		}
		for _, b := range bytes {
			stdoutChannel <- b
		}

	}
	if stdErrWriter.Len() > 0 {
		bytes := make([]byte, stdErrWriter.Len())
		_, err := stdErrWriter.Read(bytes)
		if err != nil {
			panic(err)
		}
		for _, b := range bytes {
			stdoutChannel <- b
		}
	}
}
