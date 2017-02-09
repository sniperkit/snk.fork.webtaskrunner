FROM ubuntu:16.04
#
# Install go
#
RUN \
    apt-get update && \
    apt-get install -y wget
RUN \
    wget https://storage.googleapis.com/golang/go1.7.5.linux-amd64.tar.gz -O go1.7.5.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.7.5.linux-amd64.tar.gz && \
    rm go1.7.5.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOPATH=/go

#
# Install default jdk
#
RUN \
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
RUN \
    npm install -g grunt && \
    npm install -g grunt-cli

#
# Install gulp
#
RUN \
    npm install -g gulp && \
    npm install -g gulp-cli


# Define working directory.
WORKDIR /webtaskrunner


# Define default command.
CMD ["bash"]