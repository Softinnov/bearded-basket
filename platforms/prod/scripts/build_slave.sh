#!/bin/sh

echo ">> Building db-slave image..."
cd db-slave || exit $?
docker build -t softinnov/prod-db-slave . || exit $?
cd ..
echo ">> db-slave image done."
