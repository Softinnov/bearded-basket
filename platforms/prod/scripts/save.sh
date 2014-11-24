#!/bin/sh

docker save softinnov/prod-db     > prod-db.tar || exit $?
docker save softinnov/prod-chey   > prod-chey.tar || exit $?
docker save softinnov/prod-back   > prod-back.tar || exit $?
docker save softinnov/prod-client > prod-client.tar || exit $?

zip -r prod.zip prod-db.tar prod-chey.tar prod-back.tar prod-client.tar || exit $?
