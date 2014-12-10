#!/bin/bash

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

if [ $# -ne 2  ]; then
	echo -e "$R Usage: $0 <ip> <script.sh> $W"
	exit 1
fi

ssh root@"$1" 'bash -s' < $2 || exit $?
