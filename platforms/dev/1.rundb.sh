#!/bin/bash

# RUN DB CONTAINER
# usage: ./1.rundb.sh [-t]

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

TEST=false
if [ "$1" = "-t" ]; then
	TEST=true
	shift
fi

DBDATA=dbdata
DBCON=db
DBCONTEST="$DBCON"_test

if [ $TEST = true ]; then
	echo -e "$B >> Removing old container (stop it if running) $W"
	./cleancontainer.sh $DBCONTEST

	echo -e "$B >> Running $DBCONTEST container $W"
	docker run -d --name $DBCONTEST softinnov/$DBCONTEST || exit $?
else
	echo -e "$B >> Removing old container (stop it if running) $W"
	./cleancontainer.sh $DBCON

	echo -e "$B >> Running $DBCON container $W"
	docker run -d --volumes-from $DBDATA --name $DBCON softinnov/$DBCON || exit $?
fi

echo -e "$G >> Done. $W"
