FROM golang:1.10-stretch AS builder

WORKDIR /go/src/github.com/codeclimate/hestia

ARG DEP_VERSION=v0.4.1
RUN apt-get update && \
    apt-get install -y curl git make && \
    curl -fsSL -o /usr/local/bin/dep \
    https://github.com/golang/dep/releases/download/$DEP_VERSION/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep && \
    mkdir -p /go/src/github.com/codeclimate/hestia

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only
