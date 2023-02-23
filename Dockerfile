FROM golang:1.19
SHELL ["/bin/bash", "-c"]

RUN mkdir -p /app

WORKDIR /app/

COPY . /app

EXPOSE 8080
RUN go mod download

ENTRYPOINT go run list.go
