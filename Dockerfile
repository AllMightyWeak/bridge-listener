FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o run ./cmd

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/run .

ENV WEBSOCKET=wss://
ENV PRIVATE_KEY_BRIDGE=0x_And_64symbols
ENV BRIDGE_ADDR=0x_And_40symbols
ENV WNFT_ADDR=0x_And_40symbols

CMD ["./run"]
