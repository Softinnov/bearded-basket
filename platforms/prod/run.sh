#!/bin/sh

INIT=false
USAGE="Usage: $0 [--init] <ip> <ssh_key.pub>"

if [ "$1" = "--init" ]; then
	INIT=true
	shift
	if [ $# -ne 2 ]; then
		echo $USAGE
		exit 1
	fi
else
	if [ $# -ne 2 ]; then
		echo $USAGE
		exit 1
	fi
fi


if [ $INIT = true ]; then
	./scripts/init.sh $1 $2 || exit $?

	echo ">> Constructing dbdata... from $(pwd)/data"
	./scripts/data.sh $1 || exit $?

	echo ">> Installing docker on the server..."
	./scripts/launch.sh $1 scripts/install.sh || exit $?
fi


echo ">> Building images..."
./scripts/build.sh || exit $?

echo ">> Uploading docker images on the server..."
./scripts/upload.sh $1 || exit $?

echo ">> Updating docker images on the server..."
./scripts/launch.sh $1 scripts/update.sh || exit $?
