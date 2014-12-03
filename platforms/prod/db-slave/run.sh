#!/bin/sh

INIT=false
USAGE="Usage: $0 <master ip>"
IP=$1

if [ $# -ne 1 ]; then
	echo $USAGE
	exit 1
fi

echo "\nmaster-host = $IP\n" >> /etc/mysql/conf.d/my.cnf

exec mysqld_safe
