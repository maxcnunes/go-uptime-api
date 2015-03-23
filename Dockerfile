FROM golang:1.4

WORKDIR /go/src/github.com/maxcnunes/monitor-api

ADD . /go/src/github.com/maxcnunes/monitor-api

RUN go get -d ./...

# dev environment
RUN go get github.com/codegangsta/gin \
           github.com/onsi/ginkgo/ginkgo \
           github.com/onsi/gomega
