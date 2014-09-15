#!/bin/bash

# COPY LATEST CLIENT
function cp_dir() {
rm -rf client && \
        cp -r ../../client client
	chmod -R 755 client
}

# BUILD A DOCKER CONTAINER
function docker_build() {
docker build -t softinnov/client .
}

cp_dir || exit $?
docker_build || exit $?
echo "\n>> Now you can run: $ docker run --name client --link back:back -P -d softinnov/client"

