# webtaskrunner

[![Go Report Card](https://goreportcard.com/badge/github.com/Oppodelldog/webtaskrunner)](https://goreportcard.com/report/github.com/Oppodelldog/webtaskrunner) [![GoDoc](https://godoc.org/github.com/Oppodelldog/webtaskrunner?status.svg)](https://godoc.org/github.com/Oppodelldog/webtaskrunner)

Webtaskrunner is intended to help you during development by letting you execute tasks via a webfrontend.

Set it up in your vagrant devbox or docker container and execute sometasks when needed.

####Startup
Startup the server on port :8080 with

    go run main.go
    
Then navigate to http://localhost:8080/ant

currently just ant is implemented.
The build.xml must reside in the working directory of the application.

####Add custom integrations:

* adding a go file in integrations folder, implement interface.
* add integration in main func