# Backend Dockerfile
FROM golang:1.20.6-alpine

WORKDIR /usr/app

COPY . /usr/app

RUN apk add --update make

RUN make

COPY . .

# Build the Go application
RUN go build -o mvc ./cmd/main.go

EXPOSE 8080

CMD ["./mvc"]
