#!/bin/sh

cd back || exit $?

GOOS=linux GOARCH=amd64 godep go build -o bearded-basket || exit $?

cd -

mv back/bearded-basket . || exit $?
