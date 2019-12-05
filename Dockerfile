FROM golang:1.13-alpine3.10

RUN apk update && apk add make gcc g++ git