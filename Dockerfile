# ---- Step 1: Build Stage ----
FROM golang:1.23-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go mod files separately for better caching
COPY go.mod ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application (static binary for minimal runtime dependencies)
RUN CGO_ENABLED=0 go build -o main

# ---- Step 2: Runtime Stage ----
FROM alpine:latest

# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy only the binary from the build stage
COPY --from=builder /app/main .
# Copy templates directory which contains HTML files
COPY --from=builder /app/templates ./templates

EXPOSE 8080

# Command to run the application
CMD ["./main"]
