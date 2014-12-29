#!/bin/bash

# BUILD CHEYENNE IMAGE
# usage: buildchey

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

echo -e "$B >> Building cheyenne image $W"
cd cheyenne
docker build -t softinnov/chey . || exit $?
cd ..

echo -e "$G >> Done. $W"
