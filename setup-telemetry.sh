#!/bin/bash

# Script para implementar telemetria e métricas em todos os serviços

SERVICES=("user" "payment" "notification" "catalog")
BASE_PATH="C:/programacao/backend/goexpert/go-microservices-architecture/services"

echo "🚀 Implementando telemetria e métricas em todos os serviços..."

for service in "${SERVICES[@]}"; do
    echo "📦 Processando $service..."
    
    # Criar diretórios
    mkdir -p "$BASE_PATH/$service/internal/telemetry"
    mkdir -p "$BASE_PATH/$service/internal/metrics"
    
    # Copiar arquivos de telemetria
    cp "$BASE_PATH/order/internal/telemetry/tracer.go" "$BASE_PATH/$service/internal/telemetry/"
    cp "$BASE_PATH/order/internal/metrics/metrics.go" "$BASE_PATH/$service/internal/metrics/"
    
    echo "✅ $service processado!"
done

echo "🎉 Implementação concluída!"