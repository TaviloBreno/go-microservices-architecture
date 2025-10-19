# 📊 Passo 10 - Monitoramento e Observabilidade

Este passo implementa um sistema completo de monitoramento, métricas e rastreabilidade distribuída para a arquitetura de microsserviços.

## 🎯 Componentes Implementados

### 📈 Prometheus (http://localhost:9090)
- Coleta e armazena métricas de todos os serviços
- Endpoint de métricas: `/metrics` na porta 9091 de cada serviço
- Configuração: `monitoring/prometheus.yml`

### 📊 Grafana (http://localhost:3000)
- Visualização de dashboards e alertas
- **Login**: admin / admin123
- Dashboards pré-configurados:
  - `microservices-overview.json`
  - `grpc-performance.json`
  - `business-metrics.json`

### 🔍 Jaeger (http://localhost:16686)
- Rastreamento distribuído entre serviços
- Visualização de spans e latência
- Traces completos: order → payment → notification

## 🏗️ Arquitetura de Monitoramento

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│   Microsserviços   │────▶│   Prometheus    │────▶│    Grafana      │
│   (gRPC + HTTP)   │     │   (Métricas)    │     │  (Dashboards)   │
└─────────────────┘     └─────────────────┘     └─────────────────┘
         │
         ▼
┌─────────────────┐
│     Jaeger      │
│  (Tracing)      │
└─────────────────┘
```

## 📊 Métricas Coletadas

### 🔄 Métricas Gerais (Todos os Serviços)
- `requests_total` - Total de requisições HTTP
- `http_request_duration_seconds` - Duração de requisições HTTP
- `grpc_requests_total` - Total de requisições gRPC
- `grpc_request_duration_seconds` - Duração de requisições gRPC
- `service_health` - Status de saúde do serviço

### 📦 Order Service
- `orders_processed_total` - Pedidos processados por status

### 💳 Payment Service
- `payments_processed_total` - Pagamentos processados por status e método
- `payment_amount` - Distribuição de valores de pagamento

### 🔔 Notification Service
- `notifications_sent_total` - Notificações enviadas por tipo e status
- `notification_processing_seconds` - Tempo de processamento
- `rabbitmq_queue_length` - Tamanho da fila RabbitMQ

### 👤 User Service
- `users_created_total` - Usuários criados por status

### 📋 Catalog Service
- `products_queried_total` - Consultas de produtos
- `catalog_updates_total` - Atualizações de catálogo

### 🌐 BFF GraphQL
- `graphql_queries_total` - Consultas GraphQL por operação
- `graphql_query_duration_seconds` - Duração de consultas GraphQL
- `grpc_calls_total` - Chamadas gRPC para backend services

## 🚀 Como Usar

### 1. Subir o Ambiente Completo
\`\`\`bash
docker-compose up -d --build
\`\`\`

### 2. Verificar Serviços
\`\`\`bash
docker ps
\`\`\`

### 3. Acessar Interfaces

#### Prometheus (http://localhost:9090)
- Verificar targets em **Status > Targets**
- Explorar métricas em **Graph**
- Exemplos de queries:
  \`\`\`promql
  # Taxa de requisições por serviço
  sum(rate(requests_total[5m])) by (job)
  
  # Latência P95 dos serviços gRPC
  histogram_quantile(0.95, sum(rate(grpc_request_duration_seconds_bucket[5m])) by (le, job))
  
  # Status de saúde dos serviços
  service_health
  \`\`\`

#### Grafana (http://localhost:3000)
- **Login**: admin / admin123
- Importar dashboards da pasta `monitoring/grafana/dashboards/`
- Configurar Prometheus como datasource: `http://prometheus:9090`

#### Jaeger (http://localhost:16686)
- Selecionar serviços na dropdown
- Visualizar traces distribuídos
- Analisar spans e dependências

### 4. Testar Métricas

Gere tráfego nos serviços:
\`\`\`bash
# Via BFF GraphQL (se disponível)
curl -X POST http://localhost:8080/graphql \\
  -H "Content-Type: application/json" \\
  -d '{"query": "{ orders { id status } }"}'

# Verificar métricas
curl http://localhost:9091/metrics
\`\`\`

## 🔧 Configurações Avançadas

### Endpoints de Métricas
Cada serviço expõe métricas na porta 9091:
- Order Service: http://localhost:9091/metrics
- Payment Service: http://localhost:9091/metrics  
- Notification Service: http://localhost:9091/metrics
- User Service: http://localhost:9091/metrics
- Catalog Service: http://localhost:9091/metrics
- BFF GraphQL: http://localhost:9091/metrics

### Health Checks
Todos os serviços possuem endpoint de health:
- http://localhost:9091/health

### RabbitMQ Metrics
RabbitMQ expõe métricas na porta 15692:
- http://localhost:15692/metrics

## 📈 Dashboards Disponíveis

### 1. Microservices Overview
- Requests por segundo
- Tempo de resposta
- Status de saúde dos serviços

### 2. gRPC Performance
- Taxa de requisições gRPC
- Latência P95
- Métricas específicas por serviço

### 3. Business Metrics
- Pedidos ao longo do tempo
- Taxa de sucesso de pagamentos
- Distribuição de valores
- Tempo de processamento de notificações
- Tamanho das filas RabbitMQ

## 🛠️ Instrumentação de Código

### Telemetria (OpenTelemetry)
\`\`\`go
// No main.go de cada serviço
import (
    "github.com/seu-usuario/.../internal/telemetry"
    "github.com/seu-usuario/.../internal/metrics"
)

func main() {
    // Inicializar métricas
    metrics.Init()
    
    // Inicializar tracing
    ctx := context.Background()
    shutdown := telemetry.InitTracer("service-name")
    defer shutdown(ctx)
    
    // Resto da aplicação...
}
\`\`\`

### Registrar Métricas
\`\`\`go
// Exemplo de uso nos handlers
metrics.RecordGRPCRequest("CreateOrder", "success", time.Since(start))
metrics.RecordOrderProcessed("completed")
\`\`\`

## 🎯 Próximos Passos

1. **Alertas**: Configurar alertas no Grafana baseados nas métricas
2. **SLO/SLI**: Definir Service Level Objectives e Indicators  
3. **Log Aggregation**: Integrar ELK Stack ou similar
4. **APM**: Adicionar Application Performance Monitoring
5. **Cost Monitoring**: Monitorar custos de infraestrutura

## 🚨 Troubleshooting

### Serviços não aparecem no Prometheus
1. Verificar se o serviço está expondo `/metrics` na porta 9091
2. Checar configuração em `monitoring/prometheus.yml`
3. Verificar logs: `docker logs prometheus`

### Dashboards não carregam no Grafana
1. Verificar se Prometheus está configurado como datasource
2. Importar dashboards manualmente da pasta `monitoring/grafana/dashboards/`
3. Verificar logs: `docker logs grafana`

### Traces não aparecem no Jaeger
1. Verificar se a instrumentação está ativa nos serviços
2. Checar se Jaeger está acessível nos serviços
3. Verificar logs dos serviços instrumentados

---

✅ **Sistema de Observabilidade Completo Implementado!**

O Passo 10 fornece visibilidade total da arquitetura de microsserviços com métricas, traces e dashboards profissionais para monitoramento em produção.