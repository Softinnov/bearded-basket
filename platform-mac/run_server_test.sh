#!/bin/sh

echo ">> Removing old container (stop it if running)"
./cleancontainer.sh db_test

docker run -d --name db_test softinnov/db_test

echo ">> Running the $BCONTEST container"
docker run --rm -v $GOPATH/src:/go/src --link db_test:db_test softinnov/back_test ${*:1}

echo ">> Removing old container (stop it if running)"
./cleancontainer.sh db_test
