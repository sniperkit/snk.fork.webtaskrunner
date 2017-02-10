FROM ubuntu:16.04
#
# Install go
#
RUN \
    apt-get update && \
    apt-get install -y wget && \
    apt-get install -y git && \
    wget https://storage.googleapis.com/golang/go1.7.5.linux-amd64.tar.gz -O go1.7.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.7.5.linux-amd64.tar.gz && \
    rm go1.7.5.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV PATH="/go/bin:${PATH}"
ENV GOPATH=/go

# Define working directory.
WORKDIR /webtaskrunner
RUN go get github.com/codegangsta/gin

#
# Install default jdk
#
RUN \
    apt-get update && \
    apt-get install -y default-jdk

#
# Install ant
#
RUN \
 apt-get update && \
 apt-get install -y ant


#
# Install gradle
#
RUN \
 apt-get update && \
 apt-get install -y gradle


#
# Install node
#
RUN \
    apt-get update && \
    apt-get install -y nodejs-legacy && \
    apt-get install -y npm

#
# Install grunt
#
ADD provision/package.json /node_build_tools/package.json
ADD provision/Gruntfile.js /node_build_tools/Gruntfile.js
RUN \
    npm install -g grunt && \
    npm install -g grunt-cli && \
    cd /node_build_tools && \
    npm install grunt && \
    npm install grunt-exec
#
# Install gulp
#
ADD provision/gulpfile.js /node_build_tools/gulpfile.js
RUN \
    npm install -g gulp && \
    npm install -g gulp-cli && \
    cd /node_build_tools && \
    npm install gulp && \
    npm install gulp-exec

# Define default command.
CMD ["bash"]