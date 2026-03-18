# Stage 1: build
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o api2 .

# Stage 2: runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api2 .
EXPOSE 8080
CMD ["./api2"]