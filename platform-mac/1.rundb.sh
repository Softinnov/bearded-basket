#!/bin/sh

# RUN DB CONTAINER

TEST=false
if [ "$1" == "--test" ]; then
	TEST=true
	shift
fi

DBDATA=dbdata
DBDATATEST="$DBDATA"_test
DBCON=db
DBCONTEST="$DBCON"_test

if [ $TEST = true ]; then
	echo ">> Removing old container (stop it if running)"
	./cleancontainer.sh $DBCONTEST

	echo ">> Running DB container"
	docker run -d --volumes-from $DBDATATEST --name $DBCONTEST softinnov/$DBCONTEST || exit $?
else
	echo ">> Removing old container (stop it if running)"
	./cleancontainer.sh $DBCON

	echo ">> Running DB container"
	docker run -d --volumes-from $DBDATA --name $DBCON softinnov/$DBCON || exit $?
fi
