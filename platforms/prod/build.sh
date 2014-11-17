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
rm -rf bearded-basket
cd ..
echo ">> back image done."

echo ">> Building client image..."
cd client || exit $?
RET=0
cp -r ../../client . || exit $?
docker build -t softinnov/prod-client . || RET=$?
if [ $RET -ne 0 ]; then
	rm -rf client
	exit $RET
fi
rm -rf client
cd ..
echo ">> client image done."
