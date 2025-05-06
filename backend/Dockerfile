FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o peliculas main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/peliculas .
COPY --from=builder /app/init.sql .
CMD ["./peliculas"]
