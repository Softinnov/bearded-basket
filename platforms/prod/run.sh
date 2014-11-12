#!/bin/sh

if [ $# -ne 1  ]; then
	echo "Usage: $0 <script.sh>"
	exit 1
fi

ssh root@192.99.12.123 'bash -s' < $1
