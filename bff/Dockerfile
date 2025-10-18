# ---------------------------
# Etapa de build
# ---------------------------
FROM golang:1.25 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o bff ./cmd/main.go

# ---------------------------
# Etapa final (execução)
# ---------------------------
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/bff .
EXPOSE 8080
CMD ["./bff"]
