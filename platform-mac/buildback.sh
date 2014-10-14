#!/bin/sh

# BUILD BACK IMAGE
# usage: buildback

if [ $# -ne 0 ]; then
	echo "Usage: $0"
	exit 1
fi

echo ">> Building back image"
cd back
docker build -t softinnov/dev-back . || exit $?
cd ..
