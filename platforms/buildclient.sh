#!/bin/bash

# BUILD CLIENT IMAGE
# usage: ./buildclient

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

echo -e "$B >> Building client image $W"
cd client
docker build -t softinnov/dev-client . || exit $?
cd ..

echo -e "$G >> Done. $W"
