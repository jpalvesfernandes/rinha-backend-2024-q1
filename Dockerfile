# Builder
FROM golang:1.21-alpine as builder
WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Binary
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]