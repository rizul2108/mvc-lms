# Backend Dockerfile
FROM golang:1.20.6-alpine

WORKDIR /app

# Copy only the go.mod and go.sum files to improve cache efficiency
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o mvc ./cmd/main.go

# Expose port being used
EXPOSE 9000

# Create and set the entrypoint to a shell script
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh
ENTRYPOINT ["/app/entrypoint.sh"]
