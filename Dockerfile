FROM golang

ADD . /go/src/B1210C

RUN go get golang.org/x/net/html

RUN go install B1210C

ENTRYPOINT /go/bin/B1210C

EXPOSE 8080