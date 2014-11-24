#!/bin/sh

if [ $# -ne 2  ]; then
	echo "Usage: $0 <ip> <ssh_key.pub>"
	exit 1
fi

echo ">> copy of ssh key"
ssh-copy-id -i $2 root@"$1" || exit $?

ssh root@"$1" mkdir /home/bearded-basket
