#!/bin/sh

# RUN SERVER GO
# usage: runserver <nginx> <db container> <dbuser> <dbpass> <logdir>
# example: runserver.sh localhost db admin admin $(pwd)/logs

if [ $# -ne 2 ]; then
	echo "Usage: $0 <go dir> <logdir>"
	exit 1
fi

GODIR=$1
LOGDIR=$2

echo ">> Removing old container (stop it if running)"
./cleancontainer.sh back

mkdir -p bin
cd $GODIR && \
        GOOS=linux GOARCH=amd64 go build -o bearded-basket && \
        cd - && \
        mv $GODIR/bearded-basket bin/bearded-basket

echo ">> Running the back container"
docker run --name back -v $(pwd)/bin:/exe -v $LOGDIR:/logs --link db:db -p 8002:8002 -d softinnov/dev-back

