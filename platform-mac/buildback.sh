#!/bin/sh

# BUILD BACK IMAGE
# usage: buildback

TEST=false
if [ $1 == "--test" ]; then
	TEST=true
	shift
fi

if [ $# -ne 0 ]; then
	echo "Usage: $0"
	exit 1
fi

BCON=back
BCONTEST="$BCON"_test

if [ $TEST = true ]; then
	echo ">> Building $BCONTEST image"
	cd $BCONTEST
	docker build -t softinnov/$BCONTEST . || exit $?
	cd ..
else
	echo ">> Building $BCON image"
	cd back
	docker build -t softinnov/dev-$BCON . || exit $?
	cd ..
fi
