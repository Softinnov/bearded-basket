#!/bin/sh

# RUN DB CONTAINER

echo ">> Removing old container (stop it if running)"
./cleancontainer.sh db

echo ">> Running DB container"
docker run -d --volumes-from dbdata --name db softinnov/db || exit $?

