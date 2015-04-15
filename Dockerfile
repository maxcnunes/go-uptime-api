FROM golang:1.4

WORKDIR /go/src/github.com/maxcnunes/go-uptime-api

ADD . /go/src/github.com/maxcnunes/go-uptime-api

RUN go get -d -v ./...

# dev environment
RUN go get github.com/codegangsta/gin \
           github.com/onsi/ginkgo/ginkgo \
           github.com/onsi/gomega
