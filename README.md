# webtaskrunner

quite hacky at the moment

webtaskrunner is intended to help you during development by letting you
execute tasks via a webfrontend.

set it up in your vagrant devbox or docker container and execute some
tasks when needed.

this working prototyp supports ant and it's build.xml file.

app.js contains all client logic
main.go currently contains all server logic :-)

startup server on port :8080 with

    go run main.go
    
build.xml must be in working dir