# Use official Golang image as the build stage
FROM golang:1.23-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
          CGO_ENABLED=0 \
          GOOS=linux \
          GOARCH=amd64

# Create app directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o todo-app ./cmd/main.go

# Use a minimal image for the final stage
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/todo-app .

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./todo-app"]