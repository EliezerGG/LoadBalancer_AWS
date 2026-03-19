# Stage 1: build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go mod init api2 && go build -o api2 .

# Stage 2: runtime (imagen mínima)
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api2 .
EXPOSE 3000
CMD ["./api2"]