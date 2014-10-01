#!/bin/bash

DBTABLES="toto titi tata"

if [[ $# -lt 2 ]]; then
        echo "Usage: $0 <path for cheyenne folder>"
        exit 1
fi

function test_t() {
	for DBTABLE in $DBTABLES
	do
		echo $DBTABLE
	done
}

test_t
echo $1
