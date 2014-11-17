#!/bin/sh

echo ">> Building cheyenne image..."
cd chey || exit $?
git submodule init
git submodule update
docker build -t softinnov/prod-chey . || exit $?
cd ..
echo ">> cheyenne image done."

echo ">> Building client image..."
docker build -t softinnov/prod-client . || exit $?
cd ..
echo ">> client image done."
