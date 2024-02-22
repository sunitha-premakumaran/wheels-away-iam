FROM golang:1.21.7-alpine3.19 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN apk add --no-cache gcc musl-dev
RUN go build -tags musl -o ./out/api ./cmd/api

FROM alpine:3.16.4 AS release-api
COPY --from=builder /app/out/api /app/main
COPY --from=builder /app/config /config

RUN apk update && apk add bash

CMD ["./app/main"]
