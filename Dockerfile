FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /midnight
RUN mkdir -p /midnight/proto 

WORKDIR /midnight

COPY ./proto/service.pb.go /midnight/proto
COPY ./main.go /midnight

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o midnight .

CMD ./midnight 