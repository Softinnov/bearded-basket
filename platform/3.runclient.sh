#!/bin/sh

# RUN CLIENT WEBAPP (NGINX+ANGULARJS)
# usage: runclient <webapp path> <logs dir path>

if [ $# -ne 2 ]; then
	echo "Usage: $0 <webapp> <logs dir path>"
	exit 1
fi

echo ">> Removing old container (stop it if running)"
./cleancontainer.sh client

echo ">> Running the client container"
docker run --name client -v $1:/client -v $2:/var/log/nginx --link chey:chey -p 8000:8000 -d softinnov/dev-client

