#!/bin/sh

cd ../../../server || exit $?

go get github.com/tools/godep
GOOS=linux GOARCH=amd64 godep go build -o bearded-basket || exit $?

cd -

mv ../../../server/bearded-basket . || exit $?
