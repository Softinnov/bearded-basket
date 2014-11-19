#!/bin/sh

TEST=false
if [ "$1" = "-t" ]; then
	TEST=true
	shift
else
	if [ $# -ne 4 ]; then
		echo "Usage: $0 [-t] <dbname> <dbuser> <dbpass> <dbtables>"
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
	echo ">> Removing old db_test container"
	./cleancontainer.sh db_test

	echo ">> Entering into dbtest folder"
	cd dbtests

	echo ">> Building the db_test image"
	docker build -t softinnov/db_test . || exit $?
else
	echo ">> Removing old dbdata and db container"
	./cleancontainer.sh $DBDATA
	./cleancontainer.sh $DBCON

	echo ">> Entering into db folder"
	cd db

	echo ">> Building the db image"
	docker build -t softinnov/$DBCON . || exit $?

	echo ">> Initializing the data-only container"
	docker run -d -v /var/lib/mysql --name $DBDATA softinnov/$DBCON echo data-only || exit $?

	echo ">> Initializing the mysql container"
	docker run --rm --volumes-from $DBDATA -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS softinnov/$DBCON || exit $?

	echo ">> Creating database $DBNAME for dev environment"
	docker run --rm --volumes-from $DBDATA softinnov/$DBCON bash -c "/create_db.sh $DBNAME" || exit $?

	echo ">> Importing tables $DBTABLES"
	docker run --rm -v $(pwd)/..:/data --volumes-from $DBDATA softinnov/$DBCON /bin/bash -c \
		"/import_sql.sh $DBUSER $DBPASS $DBNAME $DBTABLES" || exit $?
fi
