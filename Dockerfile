FROM ubuntu:22.04

# Backend Dockerfile
FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod vendor
RUN go mod tidy

COPY . .

# Build the Go application
RUN go build -o mvc ./cmd/main.go

FROM debian:buster-slim
COPY --from=build /app/mvc /usr/local/bin/mvc

CMD ["mvc"]
