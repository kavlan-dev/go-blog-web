FROM golang:latest AS builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app

FROM alpine:latest
WORKDIR /app/

COPY --from=builder /app/app .

CMD ["./app"]
