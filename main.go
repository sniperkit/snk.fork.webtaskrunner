package main

import (
	"net/http"
	"html/template"
	"fmt"
	"os/exec"
	"bufio"
	"bytes"
	"strings"
	"golang.org/x/net/websocket"
	"log"
	"encoding/json"
)

type Page struct {
	Title string
	Tasks []string
}

func getTaskList() []string {
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

func runCommand(cmd *exec.Cmd, errChannel chan error) {
	err := cmd.Run()
	if err != nil {
		errChannel <- err
	}

	close(errChannel)

}

func executeTask(taskName string, stdoutChannel chan string) {
	cmd := exec.Command("ant", taskName)

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

func commandHandler(ws *websocket.Conn) {
	msg := make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	taskName := string(msg[:n])
	fmt.Printf("start task: %s\n", taskName)

	outChannel := make(chan string)
	go executeTask(taskName, outChannel)

	for {
		if out, ok := <-outChannel; ok {
			if ok {
				ws.Write([]byte(out))
			}
		} else {
			break
		}
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.Handle("/cmd", websocket.Handler(commandHandler))
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{Title: "Tasks"}

	p.Tasks = getTaskList()

	t, err := template.ParseFiles("templates/index.html")
	fmt.Println(err)
	t.Execute(w, p)
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
		tasks := getTaskList()
	pJson, err := json.Marshal(tasks)
	fmt.Println(err)
	w.Header().Set("Content-Type", "application/json")
	w.Write(pJson)
}
