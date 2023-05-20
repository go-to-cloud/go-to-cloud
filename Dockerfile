## build
FROM golang:1.19-alpine3.16 AS build-env

RUN apk add build-base

ADD . /go/src/app

WORKDIR /go/src/app

RUN go env -w GO111MODULE=on \
    #&& go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o gtc

## run
FROM alpine:3.16

ADD conf/*.yaml /gotocloud/conf/

RUN mkdir -p /gotocloud && mkdir -p /gotocloud/log

WORKDIR /gotocloud

COPY --from=build-env /go/src/app/gtc /gotocloud/

ENV PATH $PATH:/gotocloud

EXPOSE 80
CMD ["/gotocloud/gtc"]