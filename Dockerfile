FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum separately to leverage Docker caching
COPY go.mod go.sum ./

# Initialize and download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the binary (Docker Buildx will set the target OS/ARCH)
RUN go build -o server .

# Create a small image for execution
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /app/server .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./server"]
