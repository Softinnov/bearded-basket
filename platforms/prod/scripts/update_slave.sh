#/bin/sh

USAGE="Usage: $0 <ip master>"

if [ $# -ne 1 ]; then
	echo $USAGE
	exit 1
fi

cd /home/bearded-basket

CNT="prod-db-slave"
ARG="docker run -d --volumes-from dbdata --name ${OLD[0]} softinnov/${OLD[0]} /run.sh $1"

echo ">> stopping and removing $CNT"
docker stop $CNT 2> /dev/null
docker rm $CNT 2> /dev/null

docker rmi softinnov/$CNT

docker load -i "$CNT".tar || exit $?

echo ">> $ARG"
$ARG || exit $?
