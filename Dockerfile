# Build stage
FROM golang:1.21-alpine AS builder

# Install git and ca-certificates (needed for fetching dependencies and HTTPS)
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create a non-root user
RUN adduser -D -g '' appuser

# Set the working directory
WORKDIR /build

# Copy go mod files first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
RUN go mod verify

# Copy the source code
COPY . .

# Build the binary with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o app .

# Final stage
FROM scratch

# Import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary from the builder stage
COPY --from=builder /build/app /app

# Copy configuration files
COPY --from=builder /build/config/config.yml /config/config.yml
COPY --from=builder /build/database/migrations /database/migrations

# Use the non-root user
USER appuser

# Expose the port (default from config)
EXPOSE 3000

# Run the binary
ENTRYPOINT ["/app"]
