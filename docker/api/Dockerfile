FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/app .
CMD ["./app"]
