FROM golang:latest
RUN mkdir -p /go/src/github.com/arturoverbel/microservice_compra
ADD . /go/src/github.com/arturoverbel/microservice_compra
WORKDIR /go/src/github.com/arturoverbel/microservice_compra
RUN go get -v
RUN go install github.com/arturoverbel/microservice_compra
ENTRYPOINT /go/bin/microservice_compra
EXPOSE 3000
