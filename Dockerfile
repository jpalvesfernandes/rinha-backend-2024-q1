# Builder
FROM golang:1.21-alpine as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Binary
FROM alpine:latest
COPY --from=builder /app/main .
CMD ["./main"]