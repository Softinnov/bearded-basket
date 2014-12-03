#!/bin/sh

if [ $# -ne 1  ]; then
	echo "Usage: $0 <ip>"
	exit 1
fi

rsync --progress -az docker-db_slave root@"$1":/home/bearded-basket/ || exit $?
