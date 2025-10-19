# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Instalar dependências do sistema
RUN apk add --no-cache git

# Copiar arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Download das dependências
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Production stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copiar o binário do builder
COPY --from=builder /app/main .

# Criar diretório para arquivos de configuração
RUN mkdir -p /root/config

# Expor a porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]