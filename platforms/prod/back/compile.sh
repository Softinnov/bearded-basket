#!/bin/sh

cd back || exit $?

go get github.com/tools/godep
godep go get ./...
GOOS=linux GOARCH=amd64 godep go build -o bearded-basket || exit $?

cd -

mv back/bearded-basket . || exit $?
