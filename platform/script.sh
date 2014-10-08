#!/bin/bash

DBNAME=prod
DBTABLES="utilisateur pdv role"
DBUSER=admin
DBPASS=admin
ANDESPATH=$1
CLIENTPATH=$2

# INIT DATABASE CONTAINER AND LAUNCH IT
function docker_db() {
	echo ">> Removing dbdata container"
	docker rm dbdata
	echo ">> Removing old container (stop it if running)"
	docker top db && docker stop db
	docker rm db
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
	echo ">> Running DB container"
	docker run -d --volumes-from dbdata --name db softinnov/db || exit $?
}

# INIT CHEYENNE CONTAINER AND LAUNCH IT
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

# INIT CLIENT CONTAINER AND LAUNCH IT
function docker_client() {
	echo ">> Building the client image"
	docker build -t softinnov/dev-client .
	echo ">> Placing logs directory"
	mkdir -p ../logs
	echo ">> Removing old container (stop it if running)"
	docker top client && docker stop client
	docker rm client
	echo ">> Running the client container"
	docker run --name client -v $CLIENTPATH:/client -v $(pwd)/../logs:/var/log/nginx --link chey:chey -p 8000:8000 -d softinnov/dev-client
}

function go_run() {
	IPNGINX=localhost
	IPDB=`docker inspect --format '{{ .NetworkSettings.IPAddress }}' db`
	go build && \
	./server -db "$DBUSER:$DBPASS@($IPDB:3306)/prod" -chey "http://$IPNGINX:8000" &> ../platform/logs/go.log &
}

if [[ $# -lt 2 ]]; then
	echo "Usage: $0 <path for cheyenne folder> <path for client folder>"
	exit 1
fi

echo ">> Moving to ./db"
cd db || exit $?
docker_db || exit $?
cd .. || exit $?
echo
echo ">> ...DB done"

echo ">> Moving to ./cheyenne"
cd cheyenne || exit $?
docker_chey || exit $?
cd .. || exit $?
echo
echo ">> ...Cheyenne done"

echo ">> Moving to ./client"
cd client || exit $?
docker_client || exit $?
cd .. || exit $?
echo
echo ">> ...Client done"

echo ">> Moving to ../server"
cd ../server || exit $?
go_run || exit $?
cd ../platform || exit $?
echo
echo ">> ...Go done"

echo
echo ">> ...everything done"
echo "You can now open your browser to http://localhost:8000"
