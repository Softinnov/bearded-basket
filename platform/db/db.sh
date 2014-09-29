#!/bin/bash

function docker_build() {
docker build -t softinnov/db .
}

docker_build || exit $?
echo "\n>> Now you can run: $ docker run --name db -d softinnov/db"

