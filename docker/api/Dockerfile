FROM golang:1.15.3-alpine3.12

COPY Gopkg.toml /go/src/horsebase-prophet/
COPY main.go /go/src/horsebase-prophet/

WORKDIR /go/src/horsebase-prophet/

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/codegangsta/gin \
  && go get -u github.com/gin-gonic/gin \
  && go get -u github.com/golang/dep/cmd/dep \
  && dep ensure

CMD gin -i run
