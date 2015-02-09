#!/bin/bash

# RUN SERVER GO
# usage: ./3.runserver [-t] <go dir> <log dir>
# example: ./3.runserver.sh $(pwd)/../server $(pwd)/logs
#
# tests example: ./3.runserver.sh -t godep go test ./...

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

TEST=false
REM=false
if [ "$1" = "-t" ]; then
	TEST=true
	shift
else
	if [ "$1" = "-tn" ]; then
		# No remove of database after tests
		REM=true
		shift
	else
		if [ $# -ne 2 ]; then
			echo -e "$R Usage: $0 [-t/-tn +commands] [<go dir> <log dir>] $W"
			exit 1
		fi
	fi
fi

GODIR=$1
LOGDIR=$2
BCON=back
BCONTEST="$BCON"_test
DBCON=db
DBCONTEST="$DBCON"_test

if [ $TEST = true ] || [ $REM = true ]; then
	./1.rundb.sh -t

	echo -e "$B >> Running the $BCONTEST container $W"
	docker run --rm -v $GOPATH/src:/go/src --link $DBCONTEST:$DBCONTEST softinnov/$BCONTEST ${*:1}

	if [ $TEST = true ]; then
		echo -e "$B >> Removing old container (stop it if running) $W"
		./cleancontainer.sh $DBCONTEST
	fi
else
	echo -e "$B >> Removing old container (stop it if running) $W"
	./cleancontainer.sh $BCON

	echo -e "$B >> Building server binary $W"
	mkdir -p bin
	cd $GODIR && \
		GOOS=linux GOARCH=amd64 godep go build -o bearded-basket && \
		cd - > /dev/null && \
		mv $GODIR/bearded-basket bin/bearded-basket

	echo -e "$B >> Running the $BCON container $W"
	docker run -v $(pwd)/bin:/exe -v $LOGDIR:/logs --link consul:consul -e SERVICE_NAME=back --name $BCON -P -d softinnov/dev-$BCON
fi

echo -e "$G >> Done. $W"
