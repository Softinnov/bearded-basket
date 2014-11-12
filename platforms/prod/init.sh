#!/bin/sh

if [ $# -ne 1  ]; then
	echo "Usage: $0 <ssh_key.pub>"
	exit 1
fi

echo ">> copy of ssh key"
ssh-copy-id -i $1 root@192.99.12.123
