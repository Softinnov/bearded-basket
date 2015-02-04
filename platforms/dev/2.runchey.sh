#!/bin/bash

# RUN CHEYENNE WEBAPP
# usage: ./2.runchey.sh <absolute path to cheyenne webapp> <log dir>
# example: ./2.runchey.sh ~/esc-pdv/src $(pwd)/logs

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

if [ $# -ne 2 ]; then
	echo -e "$R Usage: $0 <absolute path to cheyenne webapp> <log dir> $W"
	exit 1
fi

echo -e "$B >> Removing old container (stop it if running) $W"
./cleancontainer.sh chey

echo -e "$B >> Running the cheyenne container with esc-pdv path in $1 $W"
echo -e "$B >> linked with db $W"
docker run --name chey -e SERVICE_80_NAME=chey --link db:db -v $1:/esc-pdv -v $2:/var/log -d -P softinnov/chey || exit $?

echo -e "$G >> Done. $W"
