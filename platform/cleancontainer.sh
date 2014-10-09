#!/bin/sh

# STOP AND REMOVE CONTAINER BY NAME
# usage: cleancontainer.sh <container name>

if [ $# -ne 1 ]; then
	echo "Usage: $0 <container name>"
	exit 1
fi
 
RUNNING=$(docker inspect --format="{{ .State.Running }}" $1 2> /dev/null)
 
if [ $? -eq 1 ]; then
  echo "$1 does not exist."
  exit 3
fi
 
if [ "$RUNNING" = "true" ]; then
  echo ">> Stopping container: $1"
  docker stop $1
fi
 
echo ">> Clean container: $1" 
docker rm $1
