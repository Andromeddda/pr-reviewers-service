FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/prs/main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache ca-certificates postgresql-client

COPY --from=builder /app/main .

EXPOSE 8080

ENTRYPOINT ["./main"]