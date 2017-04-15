FROM golang:1.7-alpine

RUN apk update --no-cache --no-progress && \
    apk add \
    make \
    git  \
    gcc  \
    libc-dev  \
    curl


ENV LGOBIN=/go/bin
ENV PATH=$LGOBIN:$PATH

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/PI-Victor/shep

COPY glide.yaml ./ glide.lock ./

RUN glide install
