#!/bin/bash

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

if [ $# -ne 1  ]; then
	echo -e "$R Usage: $0 <ip> $W"
	exit 1
fi

echo -e "$B >> sending tar images $W"
rsync --progress -az tar_master/*.tar root@"$1":/home/bearded-basket/ || exit $?
echo -e "$G >> done. $W"
