#!/bin/sh

docker save softinnov/prod-db-slave > prod-db-slave.tar || exit $?
