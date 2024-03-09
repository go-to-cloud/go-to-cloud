## build
FROM golang:1.22.1-alpine3.19 AS build-env

RUN apk add build-base

ADD . /go/src/app

WORKDIR /go/src/app

RUN go env -w GO111MODULE=on \
    #&& go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o gtc

## run
FROM alpine:3.19

ADD conf/*.yaml /gotocloud/conf/

RUN mkdir -p /gotocloud && mkdir -p /gotocloud/log

WORKDIR /gotocloud

COPY --from=build-env /go/src/app/gtc /gotocloud/

ENV PATH $PATH:/gotocloud

EXPOSE 80
CMD ["/gotocloud/gtc"]