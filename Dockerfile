FROM golang:1.7

MAINTAINER xusss <dyxushuai@gmail.com>
ADD . /go/src/github.com/teambition/k8s-demo
RUN go get github.com/teambition/gear
RUN go install github.com/teambition/k8s-demo

ENTRYPOINT /go/bin/k8s-demo

VOLUME ["/data/log/k8s-demo"]

EXPOSE 8080
