#!/bin/sh

if [ $# -lt 2 ]; then
	echo "Usage: $0 <db data name> <db container name>"
	exit 1
fi

DBNAME=prod
DBUSER=admin
DBPASS=admin
DBDATA=$1
DBCON=$2


echo ">> Building mysql image <<$DBCON>>"
docker build -t $DBCON . || exit $?

echo ">> Initializing the volume container <<$DBDATA>>"
docker rm -v $DBDATA
docker run -d -v /var/lib/mysql --name $DBDATA $DBCON echo data-only || exit $?

echo ">> Initializing the mysql container <<$DBCON>> linked to volume <<$DBDATA>>"
docker run --rm --volumes-from $DBDATA -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS $DBCON || exit $?

echo ">> Creating database <<$DBNAME>> for dev environment"
docker run --rm --volumes-from $DBDATA $DBCON bash -c "/create_db.sh $DBNAME" || exit $?
