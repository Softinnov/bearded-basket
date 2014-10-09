#!/bin/sh

# RUN SERVER GO
# usage: runserver <cheyip> <dbip> <dbuser> <dbpass>

if [ $# -ne 4 ]; then
	echo "Usage: $0 <cheyip> <dbip> <dbuser> <dbpass>"
	exit 1
fi

echo ">> Stop it if running"
killall server

cd ../server
go build && \
	./server -db "$3:$4@($2:3306)/prod" -chey "http://$1:8000" &> ../platform/logs/go.log &
echo ">> Go server running"
