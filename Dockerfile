# Build stage
FROM golang:alpine AS build-env

# Set the working directory
WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

# Copy go.mod and go.sum files and download dependencies
COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

# Copy the entire project
COPY . .

# Build the Go application

RUN go build -o server ./src/main.go

# Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

RUN mkdir /app/logs

# Copy the binary from the build stage
COPY --from=build-env /app/server .

ENV ENV_MODE=production

# Copy configuration files
COPY production.config.yaml .

# Expose the necessary port
EXPOSE 8080

# Set the entry point
ENTRYPOINT ["./server"]