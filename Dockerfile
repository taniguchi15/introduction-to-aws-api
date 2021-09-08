FROM golang:1.17

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN apt-get update

WORKDIR /app
COPY ./ ./

RUN go get
