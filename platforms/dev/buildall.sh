#!/bin/bash

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

echo -e "$B >> Building db... $W"
./buildDB.sh prod admin admin "role utilisateur pdv"

echo -e "$B >> Building cheyenne... $W"
./buildchey.sh

echo -e "$B >> Building back... $W"
./buildback.sh

echo -e "$B >> Building client... $W"
./buildclient.sh

echo -e "$G >> Done. $W"
