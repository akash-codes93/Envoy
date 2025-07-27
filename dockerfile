# Build stage
FROM golang:1.24-alpine AS builder

# Install build tools and SQLite development headers for CGO
RUN apk add --no-cache build-base sqlite-dev

# Enable CGO so go-sqlite3 can compile
ENV CGO_ENABLED=1

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application (CGO enabled)
RUN go build -o main .

# Final stage
FROM alpine:latest

# Install runtime dependencies (SQLite libs, SSL certs, etc.)
RUN apk --no-cache add ca-certificates sqlite-libs

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]