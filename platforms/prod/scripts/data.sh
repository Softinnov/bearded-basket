#!/bin/sh

if [ $# -ne 1  ]; then
	echo "Usage: $0 <ip>"
	exit 1
fi

rm -f data.zip
zip -r data.zip data/ || exit $?

scp data.zip root@"$1":/home/bearded-basket/ || exit $?

scp -r docker-db root@"$1":/home/bearded-basket/ || exit $?
