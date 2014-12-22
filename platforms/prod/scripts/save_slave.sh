#!/bin/bash

docker save softinnov/prod-db-slave > tar_slave/prod-db-slave.tar || exit $?
