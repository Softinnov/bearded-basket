#!/bin/sh

apt-get update -y

curl -sSL https://get.docker.com/ubuntu/ | sudo sh || exit $?

docker run --rm busybox echo "everything works" || exit $?

docker run -d --volumes-from dbdata --name prod-db softinnov/prod-db || exit $?

docker run -d --link prod-db:db -v $(pwd)/logs:/var/log --name prod-chey softinnov/prod-chey || exit $?

docker run -d --link prod-db:db -v $(pwd)/logs:/logs --name prod-back softinnov/prod-back || exit $?

docker run -d --link prod-chey:chey --link prod-back:back -v $(pwd)/logs:/var/log/nginx -p 8000:8000 --name prod-client softinnov/prod-client || exit $?

