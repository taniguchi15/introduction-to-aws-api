# go build
FROM golang:1.17-alpine AS build
WORKDIR /app
ADD ./ ./
RUN CGO_ENABLED=0 go build -o build/main

# alpine image build
FROM alpine:latest AS prod
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo
WORKDIR /app
COPY --from=build /app/build/main ./
CMD ["/app/main"]

# for dev
FROM golang:1.17 AS dev

ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

WORKDIR /app
COPY ./ ./

RUN apt-get update \
  && go get
