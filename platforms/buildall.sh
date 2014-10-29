#!/bin/sh

echo ">> Building db..."
./buildDB.sh prod admin admin "role utilisateur pdv"

echo ">> Building cheyenne..."
./buildchey.sh

echo ">> Building back..."
./buildback.sh

echo ">> Building client..."
./buildclient.sh
