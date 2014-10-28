#!/bin/sh

TEST=false
if [ "$1" = "-t" ]; then
	TEST=true
	shift
fi

if [ $# -ne 4 ]; then
	echo "Usage: $0 [-t] <dbname> <dbuser> <dbpass> <dbtables>"
	exit 1
fi

DBNAME=$1
DBTEST="$DBNAME"_test
DBUSER=$2
DBPASS=$3
DBTABLES=$4
DBDATA=dbdata
DBDATATEST="$DBDATA"_test
DBCON=db
DBCONTEST="$DBCON"_test

if [ $TEST = true ]; then
	echo ">> Removing old dbdata and db container"
	./cleancontainer.sh $DBDATATEST
	./cleancontainer.sh $DBCONTEST

	echo ">> Entering db folder"
	cd db

	echo ">> Building the db image"
	docker build -t softinnov/$DBCONTEST . || exit $?

	echo ">> Initializing the data-only container"
	docker run -d -v /var/lib/mysql --name $DBDATATEST busybox echo data-only || exit $?

	echo ">> Initializing the mysql container"
	docker run --rm --volumes-from $DBDATATEST -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS softinnov/$DBCONTEST || exit $?

	echo ">> Creating database $DBTEST for test environment"
	docker run --rm --volumes-from $DBDATATEST softinnov/$DBCONTEST bash -c "/create_db.sh $DBTEST" || exit $?

	echo ">> Importing tables $DBTABLES"
	docker run --rm -v $(pwd)/..:/data --volumes-from $DBDATATEST softinnov/$DBCONTEST /bin/bash -c \
		"/import_sql.sh --test $DBUSER $DBPASS $DBTEST $DBTABLES" || exit $?
else
	echo ">> Removing old dbdata and db container"
	./cleancontainer.sh $DBDATA
	./cleancontainer.sh $DBCON

	echo ">> Entering db folder"
	cd db

	echo ">> Building the db image"
	docker build -t softinnov/$DBCON . || exit $?

	echo ">> Initializing the data-only container"
	docker run -d -v /var/lib/mysql --name $DBDATA busybox echo data-only || exit $?

	echo ">> Initializing the mysql container"
	docker run --rm --volumes-from $DBDATA -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS softinnov/$DBCON || exit $?

	echo ">> Creating database $DBNAME for dev environment"
	docker run --rm --volumes-from $DBDATA softinnov/$DBCON bash -c "/create_db.sh $DBNAME" || exit $?

	echo ">> Importing tables $DBTABLES"
	docker run --rm -v $(pwd)/..:/data --volumes-from $DBDATA softinnov/$DBCON /bin/bash -c \
		"/import_sql.sh $DBUSER $DBPASS $DBNAME $DBTABLES" || exit $?
fi
