#!/bin/sh

apt-get update -y

curl -sSL https://get.docker.com/ubuntu/ | sudo sh || exit $?

docker run --rm busybox echo "everything works" || exit $?
