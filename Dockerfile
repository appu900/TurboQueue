FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o qstash ./cmd/qstash

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/qstash .
COPY config/default.yaml ./config/default.yaml
EXPOSE 8080 9090
CMD ["./qstash"]