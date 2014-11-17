#!/bin/sh

cd back || exit $?

GOOS=linux GOARCH=amd64 go build -o bearded-basket || exit $?

cd -

mv back/bearded-basket . || exit $?
