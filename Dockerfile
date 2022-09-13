FROM golang:1.18
SHELL ["/bin/bash", "-c"]

RUN mkdir -p /app

WORKDIR /app/

COPY . /app

EXPOSE 8080
RUN go mod download

ENTRYPOINT go run srv-cal-json.go
