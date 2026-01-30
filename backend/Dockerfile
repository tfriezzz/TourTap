FROM golang:1.25-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./userapi ./cmd/server

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /build/.env ./.env
COPY --from=builder /build/userapi ./userapi
RUN chmod +x ./userapi

ENTRYPOINT ["./userapi"]

