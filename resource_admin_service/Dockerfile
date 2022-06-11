FROM golang:1.18.2

ENV KAFKA_CLUSTER_ADDR="localhost:9020"
ENV SERVER_ADDR="localhost:8080"
ENV DB_PASSWORD="xxx"

WORKDIR /app

COPY ./ /app/

RUN go mod download

ENTRYPOINT go run cmd/main.go