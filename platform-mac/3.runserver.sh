#!/bin/sh

# RUN SERVER GO
# usage: runserver <go dir> <log dir>
# example: runserver.sh $(pwd)/../server $(pwd)/logs

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
			echo "Usage: $0 [-t/-tn +commands] [<go dir> <log dir>]"
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
	./1.rundb.sh --test

	echo ">> Running the $BCONTEST container"
	docker run --rm -v $GOPATH/src:/go/src --link $DBCONTEST:$DBCONTEST softinnov/$BCONTEST ${*:1}

	if [ $TEST = true ]; then
		echo ">> Removing old container (stop it if running)"
		./cleancontainer.sh $DBCONTEST
	fi
else
	echo ">> Removing old container (stop it if running)"
	./cleancontainer.sh $BCON

	mkdir -p bin
	cd $GODIR && \
		GOOS=linux GOARCH=amd64 go build -o bearded-basket && \
		cd - > /dev/null && \
		mv $GODIR/bearded-basket bin/bearded-basket

	echo ">> Running the $BCON container"
	docker run --name $BCON -v $(pwd)/bin:/exe -v $LOGDIR:/logs --link $DBCON:$DBCON -p 8002:8002 -d softinnov/dev-$BCON
fi
