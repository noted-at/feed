FROM golang:1.26.4 AS builder
WORKDIR /app
COPY . .
RUN make xrpc
RUN CGO_ENABLED=0 GOOS=linux go build -v -tags jwx_es256k -o feed .

FROM gcr.io/distroless/static-debian13
COPY --from=builder /app/feed /feed
CMD ["/feed"]
