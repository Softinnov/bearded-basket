#!/bin/bash

# STOP AND REMOVE CONTAINER BY NAME
# usage: cleancontainer.sh <container name>

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

if [ $# -ne 1 ]; then
	echo -e "$R Usage: $0 <container name> $W"
	exit 1
fi

RUNNING=$(docker inspect --format="{{ .State.Running }}" $1 2> /dev/null)

if [ $? -eq 1 ]; then
  echo -e "$B $1 does not exist. $W"
  exit 3
fi

if [ "$RUNNING" = "true" ]; then
  echo -e "$B >> Stopping container: $1 $W"
  docker stop $1
fi

echo -e "$B >> Clean container: $1 $W"
docker rm -v $1 2> /dev/null
