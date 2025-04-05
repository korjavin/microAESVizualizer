# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./

# Copy the source code
COPY cmd/ ./cmd/

# Build the WebAssembly file
RUN GOOS=js GOARCH=wasm go build -o static/main.wasm cmd/main.go

# Build the server
RUN go build -o /app/server cmd/server/main.go

# Copy static files
COPY static/ ./static/

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the built binaries and static files from the builder stage
COPY --from=builder /app/server /app/server
COPY --from=builder /app/static /app/static

# Expose port 8080
EXPOSE 8080

# Run the server
CMD ["/app/server", "-dir", "/app/static"]