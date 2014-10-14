#!/bin/sh

if [ $# -ne 4 ]; then
	echo "Usage: $0 <dbname> <dbuser> <dbpass> <dbtables>"
	exit 1
fi

DBNAME=$1
DBUSER=$2
DBPASS=$3
DBTABLES=$4

echo ">> Removing old dbdata and db container"
./cleancontainer.sh dbdata
./cleancontainer.sh db

echo ">> Initializing the data-only container"
docker run -d -v /var/lib/mysql --name dbdata busybox echo data-only || exit $?

echo ">> Entering db folder"
cd db

echo ">> Building the db image"
docker build -t softinnov/db . || exit $?

echo ">> Initializing the mysql container"
docker run --rm --volumes-from dbdata -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS softinnov/db || exit $?

echo ">> Creating database prod"
docker run --rm --volumes-from dbdata softinnov/db bash -c "/create_db.sh prod" || exit $?

echo ">> Importing tables $DBTABLES"

docker run --rm -v $(pwd)/..:/data --volumes-from dbdata softinnov/db /bin/bash -c "/import_sql.sh $DBUSER $DBPASS $DBNAME $DBTABLES" || exit $?


