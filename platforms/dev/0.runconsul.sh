#!/bin/bash

# RUN CONSUL and REGISTRATOR containers
# usage: ./1.rundb.sh [-t]

R="\x1b[31m"
G="\x1b[32m"
B="\x1b[34m"
W="\x1b[0m"

echo -e "$B >> Removing old container (stop it if running) $W"
./cleancontainer.sh consul

echo -e "$B >> Running consul container $W"
docker run -d -p 8400:8400 -p 8500:8500 -p 172.17.42.1:53:53/udp -h consul --name consul progrium/consul -server -bootstrap -advertise 10.0.2.15 || exit $?

echo -e "$B >> Removing old container (stop it if running) $W"
./cleancontainer.sh registrator

echo -e "$B >> Running registrator container $W"
docker run -d --link consul:consul -v /var/run/docker.sock:/tmp/docker.sock --name registrator progrium/registrator consul://consul:8500

echo -e "$G >> Done. $W"
