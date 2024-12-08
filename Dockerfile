# Stage 1: Build
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o server ./cmd/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
COPY static ./static
EXPOSE 8080
CMD ["./server"]
