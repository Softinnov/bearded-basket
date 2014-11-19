#!/bin/sh

if [ $# -ne 2  ]; then
	echo "Usage: $0 <ip> <ssh_key.pub>"
	exit 1
fi

./scripts/init.sh $1 $2 || exit $?

echo ">> Installing docker on the server..."
./scripts/launch.sh $1 scripts/install.sh || exit $?

echo ">> Building images..."
./scripts/build.sh || exit $?

echo ">> Uploading docker images on the server..."
./scripts/upload.sh $1 || exit $?

echo ">> Updating docker images on the server..."
./scripts/launch.sh $1 scripts/update.sh || exit $?

