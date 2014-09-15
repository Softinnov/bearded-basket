#!/bin/bash

# BUILD SERVER BINARY
function build_bin() {
cd ../../server/ && \
        GOOS=linux GOARCH=amd64 go build bearded-basket.go && \
        cd ../platform/golang/ && \
        mv ../../server/bearded-basket .
}

# BUILD A DOCKER CONTAINER
function docker_build() {
docker build -t softinnov/server .
}

build_bin || exit $?
docker_build || exit $?
echo "\n>> Now you can run: $ docker run softinnov/server"

