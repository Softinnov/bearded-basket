#!/bin/sh

if [ $# -ne 2  ]; then
	echo "Usage: $0 <ip> <script.sh>"
	exit 1
fi

ssh root@"$1" 'bash -s' < $2
