FROM golang:1.4

WORKDIR /go/src/github.com/maxcnunes/monitor

ADD . /go/src/github.com/maxcnunes/monitor

RUN go get -d ./...