#!/bin/bash

# BUILD A DOCKER CONTAINER
function docker_build() {
docker build -t softinnov/chey .
}

docker_build || exit $?
echo 
echo ">> Now you can run: $ docker run --name chey --link db:db -v [path]:/esc-pdv -d softinnov/chey"
