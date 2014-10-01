#!/bin/bash

# BUILD A DOCKER CONTAINER
function docker_build() {
docker build -t softinnov/dev-client .
}

docker_build || exit $?
echo 
echo ">> Now you can run: $ docker run --name client -v [client]:/client --link chey:chey -P -d softinnov/dev-client"
