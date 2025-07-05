FROM golang:1.24.1-alpine AS builder

# Install git for go get, and other build tools
RUN apk add --no-cache git

WORKDIR /app

# Copy and download dependencies first (layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the source code
COPY . .

# Build the app binary
RUN go build -o imperio ./cmd

FROM alpine:3.18

# Optional: install CA certs if app uses HTTPS or DB over TLS
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy compiled binary only
COPY --from=builder /app/imperio .

# Optionally copy config/schema defaults
COPY config.yaml .
COPY schema.yaml .

# Set binary as entrypoint
ENTRYPOINT ["./imperio"]
