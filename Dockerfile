# stage 1: compile binary
FROM golang:1.21.1-alpine AS builder

WORKDIR /app

# install any necessary dependencies
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o rest-api

# stage 2: create a minimal runtime image
FROM alpine:3.18.4

WORKDIR /app

# download dockerize directly 
RUN wget -O dockerize.tar.gz https://github.com/jwilder/dockerize/releases/download/v0.7.0/dockerize-linux-amd64-v0.7.0.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize.tar.gz \
    && rm dockerize.tar.gz

# copy the binary from stage 1 to stage 2
COPY --from=builder /app/rest-api .

COPY . .
EXPOSE 8080

# use dockerize to wait for PostgreSQL and RabbitMQ to be ready before running the app
CMD ["dockerize", "-wait", "tcp://postgres:5432", "-wait", "tcp://rabbitmq:5672", "./rest-api"]
