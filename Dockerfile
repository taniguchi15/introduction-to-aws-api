FROM golang:1.17

RUN apt-get update

WORKDIR /app
ADD ["./go.mod", "./main.go", "/app"]

RUN go get
