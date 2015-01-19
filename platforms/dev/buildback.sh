#!/bin/bash

# BUILD BACK IMAGE
# usage: ./buildback [-t]

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

TEST=false
if [ "$1" = "-t" ]; then
	TEST=true
	shift
fi

BCON=back
BCONTEST="$BCON"_test

if [ $TEST = true ]; then
	echo -e "$B >> Building $BCONTEST image $W"
	cd $BCONTEST
	docker build -t softinnov/$BCONTEST . || exit $?
	cd ..
else
	echo -e "$B >> Building $BCON image $W"
	cd back
	docker build -t softinnov/dev-$BCON . || exit $?
	cd ..
fi

echo -e "$G >> Done. $W"
