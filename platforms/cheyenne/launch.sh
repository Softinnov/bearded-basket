#!/bin/bash

# BUILD A DOCKER CONTAINER
function docker_build() {
docker build -t softinnov/chey .
}

docker_build || exit $?
echo 
echo ">> Now you can run: $ docker run --name chey --link db:db -v [path]:/ANDES -d softinnov/chey"
