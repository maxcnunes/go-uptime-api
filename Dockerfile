FROM golang:1.4

WORKDIR /go/src/github.com/maxcnunes/monitor-api

ADD . /go/src/github.com/maxcnunes/monitor-api

RUN go get -d ./...

RUN go get github.com/codegangsta/gin