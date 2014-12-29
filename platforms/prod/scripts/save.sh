#!/bin/bash

docker save softinnov/prod-db         > tar_master/prod-db.tar         || exit $?
docker save softinnov/prod-esc-pdv    > tar_master/prod-esc-pdv.tar    || exit $?
docker save softinnov/prod-esc-adm    > tar_master/prod-esc-adm.tar    || exit $?
docker save softinnov/prod-esc-caisse > tar_master/prod-esc-caisse.tar || exit $?
docker save softinnov/prod-back       > tar_master/prod-back.tar       || exit $?
docker save softinnov/prod-client     > tar_master/prod-client.tar     || exit $?
