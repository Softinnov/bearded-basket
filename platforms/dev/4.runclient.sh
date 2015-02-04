#!/bin/bash

# RUN CLIENT WEBAPP (NGINX+ANGULARJS)
# usage: ./4.runclient <webapp path> <logs dir path>
# example: ./4.runclient.sh $(pwd)/../client $(pwd)/logs

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

if [ $# -ne 2 ]; then
	echo -e "$R Usage: $0 <webapp> <logs dir path> $W"
	exit 1
fi

echo -e "$B >> Removing old container (stop it if running) $W"
./cleancontainer.sh client

echo -e "$B >> Running the client container $W"
docker run --name client -v $1:/client -v $2:/var/log/nginx --link consul:consul -p 443:443 -p 8000:8000 -d softinnov/dev-client

echo -e "$G >> Done. $W"
