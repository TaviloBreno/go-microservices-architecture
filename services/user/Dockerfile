# ---------------------------
# Etapa de build
# ---------------------------
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o user-service ./cmd/main.go

# ---------------------------
# Etapa final (execução)
# ---------------------------
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/user-service .
EXPOSE 50051
CMD ["./user-service"]
