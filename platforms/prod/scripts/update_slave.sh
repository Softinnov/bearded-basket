#/bin/bash

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

USAGE="$R Usage: $0 <ip master> $W"

if [ $# -ne 1 ]; then
	echo -e $USAGE
	exit 1
fi

cd /home/bearded-basket

CNT="prod-db-slave"
ARG="docker run -d --volumes-from dbdata --name ${OLD[0]} softinnov/${OLD[0]} /run.sh $1"

echo -e "$B >> stopping and removing $CNT $W"
docker stop $CNT > /dev/null 2>&1
docker rm $CNT > /dev/null 2>&1

echo -e "$B >> removing softinnov/$CNT $W"
docker rmi softinnov/$CNT > /dev/null 2>&1

echo -e "$B >> loading "$CNT".tar $W"
docker load -i "$CNT".tar || exit $?

echo -e "$B >> $ARG $W"
$ARG || exit $?
