# Build stage
FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
COPY app.json .
#COPY start.sh .
COPY wait-for.sh .
#COPY db/migration ./db/migration


EXPOSE 8080

# Chạy start.sh trước, sau đó mới chạy main
#ENTRYPOINT ["/app/start.sh"]
CMD ["/app/main"]
