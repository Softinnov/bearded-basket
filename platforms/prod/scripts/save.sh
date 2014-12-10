#!/bin/bash

docker save softinnov/prod-db         > prod-db.tar         || exit $?
docker save softinnov/prod-esc-pdv    > prod-esc-pdv.tar    || exit $?
docker save softinnov/prod-esc-adm    > prod-esc-adm.tar    || exit $?
docker save softinnov/prod-esc-caisse > prod-esc-caisse.tar || exit $?
docker save softinnov/prod-back       > prod-back.tar       || exit $?
docker save softinnov/prod-client     > prod-client.tar     || exit $?
