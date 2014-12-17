#!/bin/sh

INIT=false
USAGE="Usage: $0 <master ip>"
IP=$1

if [ $# -ne 1 ]; then
	echo $USAGE
	exit 1
fi

exec mysqld_safe
