#!/bin/bash

####
#  This project uploads slave tarballs into the server.
####

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

if [ $# -ne 1  ]; then
	echo -e "$R Usage: $0 <ip> $W"
	exit 1
fi

IP=$1
shift

echo -e "$B >> sending tar images to $IP $W"
rsync --progress -az tar_slave/*.tar root@"$IP":/home/bearded-basket/ || exit $?
echo -e "$G >> done. $W"
