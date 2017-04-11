FROM alpine:latest

RUN apk update --no-cache --no-progress && \
    apk add go && \
    apk add make && \
    apk add bash && \
    apk add git && \
    apk add gcc && \
    apk add musl-dev && \
    apk add curl

RUN mkdir -p /go/src && \
    mkdir -p /go/bin

ENV GOPATH=/go LGOBIN=/go/bin
ENV PATH=$LGOBIN:$PATH

RUN curl https://glide.sh/get | sh

WORKDIR $GOPATH/src/github.com/PI-Victor/shep

COPY glide.yaml ./ glide.lock ./

RUN glide install
