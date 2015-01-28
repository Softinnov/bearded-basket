#!/bin/bash

####
#  This script transforms the slave mysql docker image into a tarball.
####

docker save softinnov/prod-db-slave > tar_slave/prod-db-slave.tar || exit $?
