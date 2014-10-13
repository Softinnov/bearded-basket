#!/bin/sh

# RUN SERVER GO
# usage: runserver <nginx> <db container> <dbuser> <dbpass> <logdir>
# example: runserver.sh localhost db admin admin $(pwd)/logs

if [ $# -ne 5 ]; then
	echo "Usage: $0 <cheyip> <db container> <dbuser> <dbpassa <logdir>"
	exit 1
fi

NGINX=$1
DBNAME=$2
DBUSER=$3
DBPASS=$4
LOGDIR=$5

DBIP=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' $DBNAME`


echo ">> Stop it if running"
killall server

cd ../server
go build && \
	./server -db "$DBUSER:$DBPASS@($DBIP:3306)/prod" -chey "http://$NGINX:8000" &> $LOGDIR/go.log &
echo ">> Go server running"
