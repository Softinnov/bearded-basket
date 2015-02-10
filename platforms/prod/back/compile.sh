#!/bin/sh

cd back || exit $?

go get ./...
go get github.com/tools/godep
GOOS=linux GOARCH=amd64 godep go build -o bearded-basket || exit $?

cd -

mv back/bearded-basket . || exit $?
