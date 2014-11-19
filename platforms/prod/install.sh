#!/bin/sh

apt-get update -y

curl -sSL https://get.docker.com/ubuntu/ | sudo sh || exit $?

docker run --rm busybox echo "everything works" || exit $?

PROD[0]="docker run -d --volumes-from dbdata --name prod-db softinnov/prod-db"
OLD[0]="prod-db"
PROD[1]="docker run -d --link prod-db:db -v $(pwd)/logs:/var/log --name prod-chey softinnov/prod-chey"
OLD[1]="prod-chey"
PROD[2]="docker run -d --link prod-db:db -v $(pwd)/logs:/logs --name prod-back softinnov/prod-back"
OLD[2]="prod-back"
PROD[3]="docker run -d --link prod-chey:chey --link prod-back:back -v $(pwd)/logs:/var/log/nginx -p 8000:8000 --name prod-client softinnov/prod-client"
OLD[3]="prod-client"

for i in {0..3}; do
	ARG=${PROD[$i]}
	CNT=${OLD[$i]}
	echo ">> stopping and removing $CNT"
	docker stop $CNT 2> /dev/null
	docker rm $CNT 2> /dev/null

	echo ">> $ARG"
	$ARG || exit $?
done
