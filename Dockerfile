# stage 1: compile binary
FROM golang:1.21.1-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Install any necessary dependencies
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/rest_api ./cmd/main.go

# stage 2: create a minimal runtime image
FROM alpine:3.18.4

WORKDIR /app

# Download dockerize directly
RUN wget -O dockerize.tar.gz https://github.com/jwilder/dockerize/releases/download/v0.7.0/dockerize-linux-amd64-v0.7.0.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize.tar.gz \
    && rm dockerize.tar.gz

# Copy the binary from stage 1 to stage 2
COPY --from=builder /app/rest_api .

# Copy the entire application source code
COPY . .

EXPOSE 8080

# Use dockerize to wait for PostgreSQL and RabbitMQ to be ready before running the app
CMD ["dockerize", "-wait", "tcp://postgres:5432", "-wait", "tcp://rabbitmq:5672", "./rest_api"]
