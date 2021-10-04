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

# DEPLOY/RUN container
#######################
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/bookstore .
EXPOSE 8080

CMD ["./bookstore"]

