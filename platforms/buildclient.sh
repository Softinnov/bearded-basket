#!/bin/sh

# BUILD CLIENT IMAGE
# usage: buildclient

if [ $# -ne 0 ]; then
	echo "Usage: $0"
	exit 1
fi

echo ">> Building client image"
cd client
docker build -t softinnov/dev-client . || exit $?
cd ..
