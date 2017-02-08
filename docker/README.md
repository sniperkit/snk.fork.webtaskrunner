#webtaskrunner - dev
Dockerfile

In order to work

* the following commands must be executed from "docker" subdirectory
* webtaskrunner must be located in $GOTPATH/src/github.com/Oppodelldog

###build docker image

    docker build -t"webtaskrunner-dev" .

###run webtaskrunner
    
    docker run -it  -v $(pwd)/..:/webtaskrunner -v $(echo $GOPATH)/src:/go/src/ -p 8080:8080 webtaskrunner-dev /bin/bash -c "go run main.go"
    
then open http://localhost:8080/ant


###run container in dev mode
    
    docker run -it  -v $(pwd)/..:/webtaskrunner -v $(echo $GOPATH)/src:/go/src/ webtaskrunner-dev /bin/bash
 