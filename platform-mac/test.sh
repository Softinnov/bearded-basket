#!/bin/sh

TEST=false

if [ $1 = "--test" ]; then
	TEST=true
	echo "YEAH"
	shift
fi

echo $0
echo $1
echo $TEST

if [ $TEST = true ]; then
	echo "TRUE"
else
	echo "FALSE"
fi
