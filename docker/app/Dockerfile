# Build Container
#########################
FROM golang:1.16-alpine as builder

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git bash

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY go.* ./ 
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bookstore ./cmd/bookstore/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate ./cmd/migrate/main.go

# DEPLOY/RUN container
#######################
FROM alpine:latest

RUN apk --no-cache add ca-certificates bash postgresql-client

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/bookstore .
COPY --from=builder /app/migrate .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/config.docker.toml .
COPY --from=builder /app/config.k8s-dev.toml .

COPY --from=builder /app/docker/app/bin /usr/local/bin/app
RUN chmod +x /usr/local/bin/app/
EXPOSE 8080

CMD ["./bookstore"]

