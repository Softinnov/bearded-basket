#/bin/bash

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

cd /home/bearded-basket

OLD[0]="consul"
PROD[0]="docker run -d -p 8400:8400 -p 8500:8500 -p 172.17.42.1:53:53/udp -h consul --name consul progrium/consul -server -bootstrap -advertise 10.0.2.15"
OLD[1]="registrator"
PROD[1]="docker run -d --link consul:consul -v /var/run/docker.sock:/tmp/docker.sock --name registrator progrium/registrator consul://consul:8500"
OLD[2]="db"
PROD[2]="docker run -d --volumes-from dbdata -v $(pwd)/data:/data -e SERVICE_6033_NAME=httpdb -e SERVICE_3306_NAME=db -p 6033:6033 -p 3306:3306 --name ${OLD[2]} softinnov/prod-${OLD[2]}"
OLD[3]="esc-pdv"
PROD[3]="docker run -d -e SERVICE_80_NAME=${OLD[3]} --link consul:consul -v $(pwd)/logs/pdv:/var/log --name ${OLD[3]} softinnov/prod-${OLD[3]}"
OLD[4]="esc-caisse"
PROD[4]="docker run -d -e SERVICE_80_NAME=${OLD[4]} --link consul:consul -v $(pwd)/logs/caisse:/var/log --name ${OLD[4]} softinnov/prod-${OLD[4]}"
OLD[5]="esc-adm"
PROD[5]="docker run -d -e SERVICE_80_NAME=${OLD[5]} --link consul:consul -v $(pwd)/logs/adm:/var/log --name ${OLD[5]} softinnov/prod-${OLD[5]}"
OLD[6]="back"
PROD[6]="docker run -d -e SERVICE_NAME=${OLD[6]} --link consul:consul -v $(pwd)/logs:/logs --name ${OLD[6]} softinnov/prod-${OLD[6]}"
OLD[7]="client"
PROD[7]="docker run -d --link consul:consul -v $(pwd)/logs:/var/log/nginx -v /etc/ssl/private:/etc/ssl/private -e SERVICE_80_NAME=${OLD[7]} -p 80:80 -p 443:443 --name ${OLD[7]} softinnov/prod-${OLD[7]}"

for i in {0..5}; do
	ARG=${PROD[$i]}
	CNT=${OLD[$i]}

	echo -e "$B >> stopping and removing $CNT $W"
	docker stop $CNT > /dev/null 2>&1
	docker rm $CNT > /dev/null 2>&1

	echo -e "$B >> removing softinnov/$CNT $W"
	docker rmi softinnov/$CNT > /dev/null 2>&1

	echo -e "$B >> loading "$CNT".tar $W"
	docker load -i "$CNT".tar || exit $?

	echo -e "$B >> $ARG $W"
	$ARG || exit $?
done
