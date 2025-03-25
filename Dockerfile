
# ---- Step 1: Build Stage ----
FROM golang:1.23-alpine AS builder

# Set working directory inside the container
WORKDIR /main

# Copy go mod files separately for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application (static binary for minimal runtime dependencies)
RUN CGO_ENABLED=0 go build -o ./main

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /main

EXPOSE 8080

# Command to run the application
CMD ["./main"]

