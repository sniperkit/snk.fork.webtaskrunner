#webtaskrunner - dev
Dockerfile

In order to work

* the following commands must be executed from "docker" subdirectory
* webtaskrunner must be located in $GOTPATH/src/github.com/Oppodelldog

###build docker image

    docker build -t"webtaskrunner-dev" .

###run webtaskrunner
    
    docker run -it  -v $(pwd)/..:/webtaskrunner -u $(id -u):$(id -u)  -v $(echo $GOPATH)/src:/go/src/ -p 3000:3000 webtaskrunner-dev /bin/bash -c "gin"
    
then open http://localhost:8080/ant


###run container in dev mode
    
    docker run -it  -v $(pwd)/..:/webtaskrunner -u $(id -u):$(id -u) -v $(echo $GOPATH)/src:/go/src/ webtaskrunner-dev /bin/bash
 
## assets

### build assets container image

    docker build -f assets-container -t "webtaskrunner-assets" .
    
### run assets container

    docker run -it  -v $(pwd)/../web/static:/web -u $(id -u):$(id -u) webtaskrunner-assets npm install
    


    
    