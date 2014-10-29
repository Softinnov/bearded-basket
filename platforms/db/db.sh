#!/bin/bash

function docker_build() {
docker build -t softinnov/db .
}

docker_build || exit $?
