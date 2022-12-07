FROM golang:1.17.6-alpine

RUN apk add --no-cache -X https://dl-cdn.alpinelinux.org/alpine/edge/testing watchexec
WORKDIR /go/src
COPY ./src .
RUN go get -u github.com/cosmtrek/air
CMD watchexec -w . --force-poll 100 -r go run /go/src/main.go
