#!/bin/sh

echo ">> Building db image..."
cd db || exit $?
docker build -t softinnov/prod-db . || exit $?
cd ..
echo ">> db image done."

echo ">> Building cheyenne image..."
cd chey || exit $?

ESCS="pdv adm caisse"
for E in $ESCS; do
	ESC=esc-$E
	echo ">> Fetching "$ESC"..."
	cd $E
	rm -rf $ESC && mkdir $ESC
	git archive --remote=git@bitbucket.org:softinnov/"$ESC".git --format=tar preprod | tar -xf - -C $ESC || exit $?
	cd ..
	echo ">> done."
done

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
cp -r ../../../client . || exit $?
docker build -t softinnov/prod-client . || RET=$?
if [ $RET -ne 0 ]; then
	rm -rf client
	exit $RET
fi
rm -rf client
cd ..
echo ">> client image done."
