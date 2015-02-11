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

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

USAGE="Usage: $0 [-a] [-p] [-i IMAGE]\n
  -a :\tbuild all images\n
  -p :\tpush images\n
  -i :\tprecise which image to build"

if [ -z "$1" ]; then
	echo -e $USAGE
	exit 1
fi

PUSH=false

build_esc() {
	cd chey || exit $?

	ESC=esc-$1
	echo -e "$B >> Fetching "$ESC"... $W"

	cd $1
	rm -rf $ESC && mkdir $ESC

	git archive --remote=git@bitbucket.org:softinnov/"$ESC".git --format=tar preprod | tar -xf - -C $ESC || exit $?

	echo -e "$B >> Building $ESC image... $W"
	docker build -t preprod.softinnov.fr:5000/prod-$ESC:latest . || exit $?
	cd ../..

	echo -e "$G >> done. $W"
}

build() {
	case $1 in
		db)
			echo -e "$B >> Building db image... $W"
			cd db || exit $?

			docker build -t preprod.softinnov.fr:5000/prod-db:latest . || exit $?
			cd ..

			echo -e "$G >> db image done. $W"
			;;
		esc-pdv)
			build_esc pdv
			;;
		esc-adm)
			build_esc adm
			;;
		esc-caisse)
			build_esc caisse
			;;
		back)
			echo -e "$B >> Building back image... $W"
			cd back || exit $?

			./compile.sh || exit $?
			docker build -t preprod.softinnov.fr:5000/prod-back:latest . || exit $?
			rm -rf bearded-basket
			cd ..

			echo -e "$G >> back image done. $W"
			;;
		client)
			echo -e "$B >> Building client image... $W"
			cd client || exit $?

			RET=0
			cp -r ../../../client . || exit $?

			docker build -t preprod.softinnov.fr:5000/prod-client:latest . || RET=$?
			if [ $RET -ne 0 ]; then
				rm -rf client
				exit $RET
			fi
			rm -rf client
			cd ..

			echo -e "$G >> client image done. $W"
			;;
		*)
			echo -e "$R /!\\ image $1 not found! $W"
			;;
	esac
}

push() {
	echo -e "$B >> Pushing $1 image... $W"
	docker push preprod.softinnov.fr:5000/prod-$1:latest || exit $?
	echo -e "$G >> $1 image done. $W"
}

while getopts "hai:p" opt; do
	case $opt in
		a)
			build db
			build esc-pdv
			build esc-adm
			build esc-caisse
			build back
			build client
			;;
		i)
			build $OPTARG
			;;
		p)
			PUSH=true
			;;
		[?]|h)
			echo $USAGE
			exit 1
			;;
		:)
			exit 1
			;;
	esac
done

if [ $PUSH == true ]; then
	OPTIND=1
	while getopts ":hai:p" opt; do
		case $opt in
			a)
				push db
				push esc-pdv
				push esc-adm
				push esc-caisse
				push back
				push client
				;;
			i)
				push $OPTARG
				;;
			p)
				;;
			[?]|h)
				exit 1
				;;
			:)
				exit 1
				;;
		esac
	done
fi
