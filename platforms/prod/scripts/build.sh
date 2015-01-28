#!/bin/bash

####
#  This script builds all images.
#    - softinnov/prod-db
#
#    - softinnov/prod-esc-pdv
#    - softinnov/prod-esc-adm      (it makes a git archive to pull the projects)
#    - softinnov/prod-esc-caisse
#
#    - softinnov/prod-back         (calls compile.sh to compile the last release of the server)
#
#    - softinnov/prod-client       (copies the client project here [docker build doesn't support symbolic links])
####

G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

echo -e "$B >> Building db image... $W"
cd db || exit $?
docker build -t softinnov/prod-db . || exit $?
cd ..
echo -e "$G >> db image done. $W"

echo -e "$B >> Building esc images... $W"
cd chey || exit $?

ESCS="pdv adm caisse"
for E in $ESCS; do
	ESC=esc-$E
	echo -e "$B >> Fetching "$ESC"... $W"
	cd $E
	rm -rf $ESC && mkdir $ESC
	git archive --remote=git@bitbucket.org:softinnov/"$ESC".git --format=tar preprod | tar -xf - -C $ESC || exit $?
	echo -e "$G >> done. $W"
	echo -e "$B >> Building $ESC image... $W"
	docker build -t softinnov/prod-$ESC . || exit $?
	cd ..
done

cd ..
echo -e "$G >> esc images done. $W"

echo -e "$B >> Building back image... $W"
cd back || exit $?
./compile.sh || exit $?
docker build -t softinnov/prod-back . || exit $?
rm -rf bearded-basket
cd ..
echo -e "$G >> back image done. $W"

echo -e "$B >> Building client image... $W"
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
echo -e "$G >> client image done. $W"
