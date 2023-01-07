# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /involve-postback-api

EXPOSE 3000

CMD [ "/involve-postback-api" ]
