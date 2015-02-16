#/bin/bash

####
#  This script runs all the docker containers by doing the following steps:
#
#    1. Stops the actual running container
#    2. Removes it
#    3. Removes the docker image
#    4. Loads the tarball image
#    5. Runs the container
####

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

cd /home/bearded-basket

OLD[0]="prod-db"
PROD[0]="docker run -d --volumes-from dbdata -v $(pwd)/data:/data -p 127.0.0.1:3306:3306 --name ${OLD[0]} softinnov/${OLD[0]}"
OLD[1]="prod-smtp"
PROD[1]="docker run -d -p 25:25 -e user=notification@escarcelle.net -e pass=notif512si --name smtp softinnov/${OLD[1]}"
OLD[2]="prod-esc-pdv"
PROD[2]="docker run -d --link prod-db:db --link smtp:smtp -v $(pwd)/logs/pdv:/var/log --name ${OLD[2]} softinnov/${OLD[2]}"
OLD[3]="prod-esc-caisse"
PROD[3]="docker run -d --link prod-db:db -v $(pwd)/logs/caisse:/var/log --name ${OLD[3]} softinnov/${OLD[3]}"
OLD[4]="prod-esc-adm"
PROD[4]="docker run -d --link prod-db:db -v $(pwd)/logs/adm:/var/log --name ${OLD[4]} softinnov/${OLD[4]}"
OLD[5]="prod-back"
PROD[5]="docker run -d --link prod-db:db -v $(pwd)/logs:/logs --name ${OLD[5]} softinnov/${OLD[5]}"
OLD[6]="prod-client"
PROD[6]="docker run -d --link prod-esc-pdv:esc-pdv --link prod-esc-adm:esc-adm --link prod-esc-caisse:esc-caisse --link prod-back:back -v $(pwd)/logs:/var/log/nginx -v /etc/ssl/private:/etc/ssl/private -p 80:80 -p 443:443 --name ${OLD[6]} softinnov/${OLD[6]}"

for i in {0..6}; do
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
