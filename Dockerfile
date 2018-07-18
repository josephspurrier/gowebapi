FROM golang:latest

ADD . /go/

WORKDIR /go

RUN go build app/webapi/cmd/webapi

CMD ["/go/webapi", "docker-config.json"]