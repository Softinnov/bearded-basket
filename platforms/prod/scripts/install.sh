#!/bin/sh

apt-get update -y
apt-get install -y curl

curl -sSL https://get.docker.com/ubuntu/ | sh || exit $?

docker run --rm busybox echo "everything works" || exit $?

apt-get install -y zip

cd /home/bearded-basket/

unzip data.zip || exit $?

cd docker-db

./build-db-preprod.sh dbdata db-mysql || exit $?
