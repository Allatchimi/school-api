# ----------------------- Stage building -----------------------
FROM golang:alpine AS builder
ENV LANG=C.UTF-8

WORKDIR /app

# Install packages
RUN apk update &&
apk add --update --no-cache bash make

# Install dependencies
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# Build the app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./.build/main ./cmd/main.go

# ----------------------- Stage serve -----------------------

FROM alpine:latest AS api
ENV LANG=C.UTF-8

WORKDIR /app

# Install packages
RUN apk add bash

# Copy the binary and assets to the container
COPY --from=builder /app/.build/ ./
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/.env.example ./app.env
RUN chmod +x /app/main

COPY ./docker/api/entrypoint.sh ./
RUN chmod +x /app/entrypoint.sh

EXPOSE 3100

# ENTRYPOINT ["/app/main"]
ENTRYPOINT ["/app/entrypoint.sh"]
