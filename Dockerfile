# Build stage
FROM golang:1.21-bullseye AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y gcc libsqlite3-dev && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main .

# Final stage
FROM debian:bullseye-slim

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates sqlite3 && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/config ./config
COPY --from=builder /app/db/seeds/phrases_actions.sql ./db/seeds/phrases_actions.sql

RUN mkdir -p /app/db

EXPOSE 8080

CMD sqlite3 ./db/app.db < ./db/seeds/phrases_actions.sql && ./main
