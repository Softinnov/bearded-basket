#!/bin/sh

INIT=false
USAGE="Usage: $0 [-s STEP] [--init] <ip> [ssh_key.pub]"

if [ "$1" = "--init" ]; then
	INIT=true
	shift
	if [ $# -ne 2 ]; then
		echo $USAGE
		exit 1
	fi
else
	if [ $# -ne 1 ]; then
		echo $USAGE
		exit 1
	fi
fi


if [ $INIT = true ]; then
	echo "\n======= STEP 1 =========\n"
	./scripts/init.sh $1 $2 || exit $?

	echo "\n======= STEP 2 =========\n"
	echo ">> Constructing dbdata... from $(pwd)/data"
	./scripts/data.sh $1 || exit $?

	echo "\n======= STEP 3 =========\n"
	echo ">> Installing docker on the server..."
	./scripts/launch.sh $1 scripts/install.sh || exit $?
fi

echo "\n======= STEP 4 =========\n"
echo ">> Building images..."
./scripts/build.sh || exit $?

echo "\n======= STEP 5 =========\n"
echo ">> Save images into zipfile"
./scripts/save.sh || exit $?

echo "\n======= STEP 6 =========\n"
echo ">> Uploading docker images on the server..."
./scripts/upload.sh $1 || exit $?

echo "\n======= STEP 7 =========\n"
echo ">> Updating docker images on the server..."
./scripts/launch.sh $1 scripts/update.sh || exit $?
