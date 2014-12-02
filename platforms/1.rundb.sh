#!/bin/sh

# RUN DB CONTAINER

TEST=false
if [ "$1" = "-t" ]; then
	TEST=true
	shift
fi

DBDATA=dbdata
DBCON=db
DBCONTEST="$DBCON"_test

if [ $TEST = true ]; then
	echo ">> Removing old container (stop it if running)"
	./cleancontainer.sh $DBCONTEST

	echo ">> Running $DBCONTEST container"
	docker run -d --name $DBCONTEST softinnov/$DBCONTEST || exit $?
else
	echo ">> Removing old container (stop it if running)"
	./cleancontainer.sh $DBCON

	echo ">> Running $DBCON container"
	docker run -d --volumes-from $DBDATA --name $DBCON softinnov/$DBCON || exit $?
fi
