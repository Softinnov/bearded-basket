
FROM golang:1.3
MAINTAINER Ghislain Guiot

RUN echo Europe/Paris > /etc/timezone && dpkg-reconfigure --frontend noninteractive tzdata

RUN go get code.google.com/p/go.tools/cmd/cover
RUN go get github.com/tools/godep

# SHARE $GOPATH/src for tests
VOLUME /go/src

WORKDIR /go/src/github.com/Softinnov/bearded-basket/server
