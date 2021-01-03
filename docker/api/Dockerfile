FROM golang:1.15.3-alpine3.12

WORKDIR /go/src/api/

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/gin-gonic/gin \
  && go get -u github.com/jinzhu/gorm \
  && go get -u github.com/go-sql-driver/mysql \
  && go get -u github.com/derekparker/delve/cmd/dlv

