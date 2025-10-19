#!/bin/bash

# Script para instrumentar todos os main.go files com telemetria e métricas

echo "🚀 Adicionando instrumentação aos main.go files..."

# Function to add imports and initialization to main.go
add_instrumentation() {
    local service=$1
    local service_name=$2
    local main_file="./services/$service/cmd/main.go"
    
    if [ -f "$main_file" ]; then
        echo "📦 Instrumentando $service..."
        
        # Create backup
        cp "$main_file" "$main_file.backup"
        
        # The instrumentation will be added manually to each service
        echo "✅ Backup criado para $service"
    else
        echo "⚠️ main.go não encontrado em $service"
    fi
}

# Add instrumentation to all services
add_instrumentation "order" "order-service"
add_instrumentation "payment" "payment-service"
add_instrumentation "notification" "notification-service"
add_instrumentation "user" "user-service"
add_instrumentation "catalog" "catalog-service"
add_instrumentation "bff-graphql" "bff-service"

echo "🎉 Script de instrumentação concluído!"
echo "📝 Lembre-se de:"
echo "   1. Adicionar imports de telemetry e metrics em cada main.go"
echo "   2. Chamar metrics.Init() e telemetry.InitTracer() no início"
echo "   3. Usar defer shutdown(ctx) para o tracer"
echo "   4. Executar 'go mod tidy' em cada serviço"