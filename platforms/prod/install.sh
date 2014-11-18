#!/bin/sh

apt-get update -y

curl -sSL https://get.docker.com/ubuntu/ | sudo sh || exit $?

docker run --rm busybox echo "everything works" || exit $?

docker run -d --volumes-from dbdata --name db softinnov/prod-db || exit $?

docker run -d --link db:db -v $(pwd)/logs:/var/log --name chey softinnov/prod-chey || exit $?

docker run -d --link db:db -v $(pwd)/logs:/logs --name back softinnov/prod-back || exit $?

docker run -d -v $(pwd)/logs:/var/log/nginx --link chey:chey --link back:back -p 8000:8000 --name client softinnov/prod-client || exit $?

