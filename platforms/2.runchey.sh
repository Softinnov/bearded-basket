#!/bin/sh

# RUN CHEYENNE WEBAPP
# usage: runchey <absolute path to cheyenne webapp> <log dir>

if [ $# -ne 2 ]; then
	echo "Usage: $0 <absolute path to cheyenne webapp> <log dir>"
	exit 1
fi

echo ">> Removing old container (stop it if running)"
./cleancontainer.sh chey

echo ">> Running the cheyenne container with esc-pdv path in $1"
echo ">> linked with db"
docker run --name chey --link db:db -v $1:/esc-pdv -v $2:/var/log -d -p 8001:80 softinnov/chey || exit $?

