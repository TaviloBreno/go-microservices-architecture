# ---------------------------
# Etapa de build
# ---------------------------
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o order-service ./cmd/main.go

# ---------------------------
# Etapa final (execução)
# ---------------------------
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/order-service .
EXPOSE 50053
CMD ["./order-service"]
