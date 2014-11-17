#!/bin/sh

echo ">> Building db image..."
cd db || exit $?
docker build -t softinnov/prod-db . || exit $?
cd ..
echo ">> db image done."

echo ">> Building cheyenne image..."
cd chey || exit $?
git submodule init
git submodule update
docker build -t softinnov/prod-chey . || exit $?
cd ..
echo ">> cheyenne image done."

echo ">> Building back image..."
cd back || exit $?
./compile.sh || exit $?
docker build -t softinnov/prod-back . || exit $?
cd ..
echo ">> back image done."

echo ">> Building client image..."
cd client || exit $?
docker build -t softinnov/prod-client . || exit $?
cd ..
echo ">> client image done."

