#!/bin/sh

# BUILD CHEYENNE IMAGE
# usage: buildchey

if [ $# -ne 0 ]; then
	echo "Usage: $0"
	exit 1
fi

echo ">> Building cheyenne image"
cd cheyenne
docker build -t softinnov/chey . || exit $?
cd ..
