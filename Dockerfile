FROM golang:1.25.1 AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -tags jwx_es256k -o feed .

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/feed /feed
CMD ["/feed"]
