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
RUN go build -o be-web ./src/main.go

# Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

RUN mkdir /app/logs

# Copy the binary from the build stage
COPY --from=build-env /app/be-web .

# Copy configuration files
COPY dev.config.yaml .

# Expose the necessary port
EXPOSE 8080

# Set the entry point
ENTRYPOINT ["./be-web"]