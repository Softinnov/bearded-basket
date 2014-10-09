#!/bin/sh

# RUN CHEYENNE WEBAPP
# usage: runchey <absolute path to cheyenne webapp>

if [ $# -ne 1 ]; then
	echo "Usage: $0 <absolute path to cheyenne webapp>"
	exit 1
fi

echo ">> Removing old container (stop it if running)"
./cleancontainer.sh chey

echo ">> Running the cheyenne container with ANDES path in $1"
echo ">> linked with db"
docker run --name chey --link db:db -v $1:/ANDES -d softinnov/chey || exit $?

