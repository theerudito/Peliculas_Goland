FROM golang:1.24-alpine AS builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY . .
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o peliculas main.go
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/peliculas .
COPY --from=builder /app/init.sql .
CMD ["./peliculas"]
