#/bin/sh

cd /home/bearded-basket

OLD[0]="prod-db"
PROD[0]="docker run -d --volumes-from dbdata -v $(pwd)/data:/data --name ${OLD[0]} softinnov/${OLD[0]}"

# Only run the docker "prod-db-slave" with run.sh <ip master>

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
