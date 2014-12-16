#!/bin/bash

G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

echo -e "$B >> installing docker $W"
curl -sSL https://get.docker.com/ubuntu/ | sh || exit $?
echo -e "$G >> done. $W"

echo -e "$B >> test of docker $W"
docker run --rm busybox echo -e "everything works" || exit $?
echo -e "$G >> done. $W"

cd /home/bearded-basket/docker-db_slave

echo -e "$B >> build db-preprod $W"
./build-db-preprod.sh dbdata db-mysql || exit $?
echo -e "$G >> done. $W"
