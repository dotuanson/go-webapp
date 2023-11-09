# Build stage
FROM golang:1.20-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go


# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .

RUN addgroup -S app-group && adduser -S app-user -G app-group
USER app-user

EXPOSE 8000
CMD ["./main"]