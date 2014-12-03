#/bin/sh

cd /home/bearded-basket

rm -rf prod-*.tar
unzip prod.zip || exit $?

OLD[0]="prod-db"
PROD[0]="docker run -d --volumes-from dbdata -v $(pwd)/data:/data --name ${OLD[0]} softinnov/${OLD[0]}"
OLD[1]="prod-esc-pdv"
PROD[1]="docker run -d --link prod-db:db -v $(pwd)/logs/pdv:/var/log --name ${OLD[1]} softinnov/${OLD[1]}"
OLD[2]="prod-esc-caisse"
PROD[2]="docker run -d --link prod-db:db -v $(pwd)/logs/caisse:/var/log --name ${OLD[2]} softinnov/${OLD[2]}"
OLD[3]="prod-esc-adm"
PROD[3]="docker run -d --link prod-db:db -v $(pwd)/logs/adm:/var/log --name ${OLD[3]} softinnov/${OLD[3]}"
OLD[4]="prod-back"
PROD[4]="docker run -d --link prod-db:db -v $(pwd)/logs:/logs --name ${OLD[4]} softinnov/${OLD[4]}"
OLD[5]="prod-client"
PROD[5]="docker run -d --link prod-esc-pdv:esc-pdv --link prod-esc-adm:esc-adm --link prod-esc-caisse:esc-caisse --link prod-back:back -v $(pwd)/logs:/var/log/nginx -v /etc/ssl/private:/etc/ssl/private -p 80:8000 -p 443:443 --name ${OLD[5]} softinnov/${OLD[5]}"

for i in {0..5}; do
	ARG=${PROD[$i]}
	CNT=${OLD[$i]}

	echo ">> stopping and removing $CNT"
	docker stop $CNT 2> /dev/null
	docker rm $CNT 2> /dev/null

	docker rmi softinnov/$CNT

	docker load -i "$CNT".tar || exit $?

	echo ">> $ARG"
	$ARG || exit $?
done
