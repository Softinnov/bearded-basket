#!/bin/bash

docker save softinnov/prod-db-slave > prod-db-slave.tar || exit $?
