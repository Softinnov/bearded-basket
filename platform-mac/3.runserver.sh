#!/bin/sh

# RUN SERVER GO
# usage: runserver <nginx> <db container> <dbuser> <dbpass> <logdir>
# example: runserver.sh localhost db admin admin $(pwd)/logs

TEST=false
if [ "$1" == "--test" ]; then
	TEST=true
	shift
else
	if [ $# -ne 2 ]; then
		echo "Usage: $0 [--test +commands] [<go dir> <logdir>]"
		exit 1
	fi
fi

GODIR=$1
LOGDIR=$2
BCON=back
BCONTEST="$BCON"_test
DBCON=db
DBCONTEST="$DBCON"_test

if [ $TEST = true ]; then
	echo ">> Removing old container (stop it if running)"
	./cleancontainer.sh $BCONTEST

	echo ">> Running the $BCONTEST container"
	docker run -v $GOPATH/src:/go/src --link $DBCONTEST:$DBCONTEST softinnov/$BCONTEST ${*:1}
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
