# webtaskrunner

[![Go Report Card](https://goreportcard.com/badge/github.com/Oppodelldog/webtaskrunner)](https://goreportcard.com/report/github.com/Oppodelldog/webtaskrunner) [![GoDoc](https://godoc.org/github.com/Oppodelldog/webtaskrunner?status.svg)](https://godoc.org/github.com/Oppodelldog/webtaskrunner)

Webtaskrunner is intended to help you during development by letting you execute tasks via a webfrontend.

Set it up in your vagrant devbox or docker container and execute some tasks when needed.

####Startup
Startup the server on port :8080 with

    go run main.go
    
Then navigate to http://localhost:8080/ant

Currently **ant**, **gradle**, **grunt** and **gulp** are integrated. 
The appropriate build files must reside in the working directory of the application.
For gulp and grunt the location of the buildfile can be configured in webtaskrunner.yaml.


####Add custom integrations:
Is there a build tool, that is yet not supported?

* add a go file in integrations folder, implement the 'integration' interface.
* add integration in main func