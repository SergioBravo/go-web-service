FROM golang:1.15.3-alpine

WORKDIR /opt/code
ADD ./ /opt/code

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go mod download

RUN go build -o bin/go-web-service cmd/webservice/main.go
ENTRYPOINT [ "bin/go-web-service" ]