#!/bin/bash

# BUILD DB IMAGE
# usage: ./buildDB [-t] <dbname> <dbuser> <dbpass> <dbtables>
# example: /buildDB.sh prod admin admin "role utilisateur pdv"

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

TEST=false
if [ "$1" = "-t" ]; then
	TEST=true
	shift
else
	if [ $# -ne 4 ]; then
		echo -e "$R Usage: $0 [-t] <dbname> <dbuser> <dbpass> <dbtables> $W"
		exit 1
	fi
fi

DBNAME=$1
DBUSER=$2
DBPASS=$3
DBTABLES=$4
DBDATA=dbdata
DBCON=db

if [ $TEST = true ]; then
	echo -e "$B >> Removing old db_test container $W"
	./cleancontainer.sh db_test

	cd dbtests

	echo -e "$B >> Building the db_test image $W"
	docker build -t softinnov/db_test . || exit $?
else
	echo -e "$B >> Removing old dbdata and db container $W"
	./cleancontainer.sh $DBDATA
	./cleancontainer.sh $DBCON

	cd db

	echo -e "$B >> Building the db image $W"
	docker build -t softinnov/$DBCON . || exit $?

	echo -e "$B >> Initializing the data-only container $W"
	docker run -d -v /var/lib/mysql --name $DBDATA softinnov/$DBCON echo data-only || exit $?

	echo -e "$B >> Initializing the mysql container $W"
	docker run --rm --volumes-from $DBDATA -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS softinnov/$DBCON || exit $?

	echo -e "$B >> Creating database $DBNAME for dev environment $W"
	docker run --rm --volumes-from $DBDATA softinnov/$DBCON bash -c "/create_db.sh $DBNAME" || exit $?

	echo -e "$B >> Importing tables $DBTABLES $W"
	docker run --rm -v $(pwd)/..:/data --volumes-from $DBDATA softinnov/$DBCON /bin/bash -c \
		"/import_sql.sh $DBUSER $DBPASS $DBNAME $DBTABLES" || exit $?
fi

echo -e "$G >> Done. $W"
