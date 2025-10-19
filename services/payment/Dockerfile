# Multi-stage build
FROM golang:1.25-alpine AS builder

# Instalar dependências de build
RUN apk add --no-cache git

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Imagem final
FROM alpine:latest

# Instalar ca-certificates para requisições HTTPS
RUN apk --no-cache add ca-certificates

# Criar usuário não-root
RUN adduser -D -s /bin/sh appuser

# Definir diretório de trabalho
WORKDIR /root/

# Copiar o binário da aplicação
COPY --from=builder /app/main .

# Mudar para usuário não-root
USER appuser

# Expor porta
EXPOSE 50053

# Comando para executar a aplicação
CMD ["./main"]