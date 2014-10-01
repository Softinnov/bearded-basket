#!/bin/bash

DBNAME=prod
DBTABLES="utilisateur pdv"
DBUSER=admin
DBPASS=admin
ANDESPATH=$1

# INIT DATABASE CONTAINER AND LAUNCH IT
function docker_db() {
	echo ">> Removing dbdata container"
	docker rm dbdata
	echo ">> Initializing the data-only container"
	docker run -d -v /var/lib/mysql --name dbdata busybox echo data-only || exit $?
	echo ">> Building the db image"
	docker build -t softinnov/db . || exit $?
	echo ">> Initializing the mysql container"
	docker run --rm --volumes-from dbdata -e MYSQL_USER=$DBUSER -e MYSQL_PASS=$DBPASS softinnov/db || exit $?
	echo ">> Creating database prod"
	docker run --rm --volumes-from dbdata softinnov/db bash -c "/create_db.sh prod" || exit $?
	echo ">> Importing tables $DBTABLES"
	cd .. && \
	docker run --rm -v $(pwd):/data --volumes-from dbdata softinnov/db /bin/bash -c "/import_sql.sh $DBUSER $DBPASS $DBNAME $DBTABLES" || exit $?
	cd db
	echo ">> Removing old container (stop it if running)"
	docker top db && docker stop db
	docker rm db
	echo ">> Running DB container"
	docker run -d --volumes-from dbdata --name db softinnov/db || exit $?
}

function docker_chey() {
	echo ">> Building the Cheyenne image"
	docker build -t softinnov/chey . || exit $?
	echo ">> Removing old container (stop it if running)"
	docker top chey && docker stop chey
	docker rm chey
	echo ">> Running the cheyenne container with ANDES path in $ANDESPATH"
	echo ">> linked with db"
	docker run --name chey --link db:db -v $ANDESPATH:/ANDES -d softinnov/chey || exit $?
}

function docker_go() {
}

if [[ $# -lt 1 ]]; then
	echo "Usage: $0 <path for cheyenne folder>"
	exit 1
fi

# echo ">> Moving to ./db"
# cd db || exit $?
# docker_db || exit $?
# cd .. || exit $?
# echo
# echo ">> ...DB done"

# echo ">> Moving to ./cheyenne"
# cd cheyenne || exit $?
# docker_chey || exit $?
# cd .. || exit $?
# echo
# echo ">> ...Cheyenne done"

echo ">> Moving to ./server"
cd server || exit $?
docker_go || exit $?
cd .. || exit $?
echo
echo ">> ...Go backend done"

echo
echo ">> ...everything done"
