#!/bin/bash

####
#  This script runs all the docker containers by doing the following steps:
#
#    1. Stops the actual running container
#    2. Removes it
#    3. Pulls the docker image
#    4. Runs the container
####

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

cd /home/bearded-basket

OLD[0]="consul"
PROD[0]="docker run -d -p 8400:8400 -p 8500:8500 -p 172.17.42.1:53:53/udp -h consul --name consul progrium/consul -server -bootstrap -advertise 172.17.42.1"
OLD[1]="registrator"
PROD[1]="docker run -d --link consul:consul -v /var/run/docker.sock:/tmp/docker.sock --name registrator progrium/registrator consul://consul:8500"
OLD[2]="db"
PROD[2]="docker run -d -e SERVICE_6033_NAME=httpdb -e SERVICE_3306_NAME=db --volumes-from dbdata -v $(pwd)/data:/data -p 6033:6033 -p 3306:3306 --name ${OLD[2]} preprod.softinnov.fr:5000/prod-${OLD[2]}"
OLD[3]="esc-pdv"
PROD[3]="docker run -d -e SERVICE_80_NAME=${OLD[3]} --link consul:consul -v $(pwd)/logs/pdv:/var/log    -P --name ${OLD[3]} preprod.softinnov.fr:5000/prod-${OLD[3]}"
OLD[4]="esc-caisse"
PROD[4]="docker run -d -e SERVICE_80_NAME=${OLD[4]} --link consul:consul -v $(pwd)/logs/caisse:/var/log -P --name ${OLD[4]} preprod.softinnov.fr:5000/prod-${OLD[4]}"
OLD[5]="esc-adm"
PROD[5]="docker run -d -e SERVICE_80_NAME=${OLD[5]} --link consul:consul -v $(pwd)/logs/adm:/var/log    -p 127.0.0.1::80 --name ${OLD[5]} preprod.softinnov.fr:5000/prod-${OLD[5]}"
OLD[6]="back"
PROD[6]="docker run -d -e SERVICE_NAME=${OLD[6]}    --link consul:consul -v $(pwd)/logs:/logs           -p 127.0.0.1::8002 --name ${OLD[6]} preprod.softinnov.fr:5000/prod-${OLD[6]}"
OLD[7]="client"
PROD[7]="docker run -d -e SERVICE_80_NAME=${OLD[7]} --link consul:consul -v $(pwd)/logs:/var/log/nginx -v /etc/ssl/private:/etc/ssl/private -p 80:80 -p 443:443 --name ${OLD[7]} preprod.softinnov.fr:5000/prod-${OLD[7]}"

for i in {0..7}; do
	ARG=${PROD[$i]}
	CNT=${OLD[$i]}

	echo -e "$B >> stopping and removing $CNT $W"
	docker stop $CNT > /dev/null 2>&1
	docker rm $CNT > /dev/null 2>&1

	if [ $i -ge 2 ]; then
		echo -e "$B >> removing preprod.softinnov.fr:5000/prod-$CNT $W"
		docker pull preprod.softinnov.fr:5000/prod-$CNT || exit $?
	else
		echo -e "$B >> removing progrium/$CNT $W"
		docker pull progrium/$CNT || exit $?
	fi


	echo -e "$B >> $ARG $W"
	$ARG || exit $?
done
