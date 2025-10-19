# 🚀 Guia de Início Rápido

Este guia vai te ajudar a colocar toda a arquitetura de microserviços rodando em poucos minutos.

## 📋 Pré-requisitos

Antes de começar, certifique-se de ter instalado:

- **Docker** (versão 20.10+) e **Docker Compose** (versão 2.0+)
- **Go** (versão 1.21+)
- **Node.js** (versão 18+) e **npm**
- **Git**

### Verificar instalações

```bash
docker --version
docker-compose --version
go version
node --version
npm --version
```

## 🏃‍♂️ Início Rápido (3 minutos)

### 1. Clone o repositório

```bash
git clone <repository-url>
cd go-microservices-architecture
```

### 2. Configure as variáveis de ambiente (opcional)

As configurações padrão já funcionam, mas você pode customizar:

```bash
# Copie o arquivo de exemplo (se existir)
cp .env.example .env

# Edite conforme necessário
nano .env
```

### 3. Suba toda a infraestrutura

```bash
# Subir todos os serviços de uma vez
docker-compose up -d

# Ou use o Makefile (mais fácil)
make docker-up
```

### 4. Aguarde os serviços iniciarem

```bash
# Acompanhe os logs
docker-compose logs -f

# Ou verifique o status
make docker-ps
```

### 5. Verifique a saúde dos serviços

```bash
# Execute o health check
bash scripts/health-check.sh

# Ou use o Makefile
make health-check
```

### 6. Acesse as interfaces

Abra seu navegador e acesse:

| Serviço | URL | Credenciais |
|---------|-----|-------------|
| 🎨 **Frontend React** | http://localhost:3001 | - |
| 🔷 **GraphQL Playground** | http://localhost:8080/graphql | - |
| 📊 **Grafana** | http://localhost:3000 | admin / admin123 |
| 📈 **Prometheus** | http://localhost:9090 | - |
| 🔍 **Jaeger** | http://localhost:16686 | - |
| 🐰 **RabbitMQ** | http://localhost:15672 | guest / guest |

## 🎯 Testando a Aplicação

### 1. Via Interface Gráfica

Acesse o frontend em http://localhost:3001 e navegue pelas funcionalidades.

### 2. Via GraphQL Playground

Acesse http://localhost:8080/graphql e experimente as queries:

```graphql
# Criar um usuário
mutation {
  createUser(input: {
    name: "João Silva"
    email: "joao@example.com"
    password: "senha123"
  }) {
    id
    name
    email
  }
}

# Listar usuários
query {
  users {
    id
    name
    email
    createdAt
  }
}

# Criar um pedido
mutation {
  createOrder(input: {
    userId: "1"
    items: [
      { productId: "1", quantity: 2, price: 99.90 }
    ]
    totalAmount: 199.80
  }) {
    id
    userId
    totalAmount
    status
  }
}

# Processar pagamento
mutation {
  processPayment(input: {
    orderId: "1"
    amount: 199.80
    paymentMethod: "credit_card"
    cardNumber: "4111111111111111"
  }) {
    id
    orderId
    status
    amount
  }
}
```

### 3. Via cURL

```bash
# Health Check do Order Service
curl http://localhost:50051/health

# Métricas do Order Service
curl http://localhost:50051/metrics

# GraphQL Query
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query":"{ users { id name email } }"}'
```

## 📊 Monitoramento

### Grafana Dashboards

1. Acesse http://localhost:3000 (admin/admin123)
2. Navegue para **Dashboards**
3. Explore os dashboards pré-configurados:
   - **Microservices Overview** - Visão geral de todos os serviços
   - **gRPC Performance** - Performance das chamadas gRPC
   - **Business Metrics** - Métricas de negócio

### Prometheus Queries

Acesse http://localhost:9090 e experimente:

```promql
# Taxa de requisições por segundo
rate(grpc_server_handled_total[5m])

# Latência P95
histogram_quantile(0.95, rate(grpc_server_handling_seconds_bucket[5m]))

# Total de pedidos criados
orders_created_total

# Pagamentos por método
payments_processed_total
```

### Jaeger Tracing

1. Acesse http://localhost:16686
2. Selecione um serviço (ex: order-service)
3. Clique em **Find Traces**
4. Explore a latência e dependências entre serviços

## 🧪 Executando Testes

### Todos os testes

```bash
# Via script
bash scripts/run-all-tests.sh

# Ou via Makefile
make test

# Ou simule o CI localmente
make ci-local
```

### Testes de um serviço específico

```bash
# Order Service
make test-order

# Payment Service
make test-payment

# User Service
make test-user

# Frontend
make test-frontend
```

### Cobertura de código

```bash
# Gerar relatório de cobertura
make coverage

# Ver cobertura detalhada de um serviço
cd services/order
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 🛠️ Desenvolvimento

### Rodando um serviço individualmente

```bash
# Order Service
make dev-order

# Payment Service
make dev-payment

# User Service
make dev-user

# BFF GraphQL
make dev-bff

# Frontend
make dev-frontend
```

### Rodando linter

```bash
# Checar código
make lint

# Auto-fix problemas
make lint-fix
```

### Build local

```bash
# Buildar todos os serviços
make build

# Buildar apenas o frontend
make build-frontend
```

## 🐳 Comandos Docker Úteis

```bash
# Parar todos os containers
make docker-down

# Reiniciar tudo
make docker-restart

# Ver logs de todos os serviços
make docker-logs

# Ver logs de um serviço específico
docker-compose logs -f order-service

# Rebuild imagens
make docker-build

# Limpar tudo (volumes inclusos)
make clean-docker
```

## 🗄️ Database

### Reset do banco de dados

```bash
# Resetar banco (cuidado em produção!)
make db-reset
```

### Acessar MySQL

```bash
# Via Docker
docker-compose exec mysql mysql -u root -p

# Senha: root123

# Usar o banco
USE microservices;

# Ver tabelas
SHOW TABLES;

# Consultar dados
SELECT * FROM orders LIMIT 10;
```

## 🚨 Troubleshooting

### Problema: Porta já em uso

```bash
# Identificar processo usando a porta
# Windows
netstat -ano | findstr :3306

# Linux/Mac
lsof -i :3306

# Matar processo
# Windows
taskkill /PID <PID> /F

# Linux/Mac
kill -9 <PID>
```

### Problema: Container não inicia

```bash
# Ver logs detalhados
docker-compose logs order-service

# Recriar container
docker-compose up -d --force-recreate order-service

# Verificar recursos do Docker
docker system df
```

### Problema: Testes falhando

```bash
# Limpar cache de testes
go clean -testcache

# Rodar testes com verbose
cd services/order
go test ./... -v

# Testar pacote específico
go test ./internal/service -v
```

### Problema: Frontend não carrega

```bash
# Limpar node_modules e reinstalar
cd frontend
rm -rf node_modules package-lock.json
npm install

# Limpar cache do npm
npm cache clean --force
```

## 📁 Estrutura de Pastas

```
.
├── .github/
│   └── workflows/          # CI/CD pipelines
├── bff/                    # Backend for Frontend (GraphQL)
├── data/                   # Dados persistentes
├── docs/                   # Documentação
├── frontend/               # React Dashboard
├── infra/                  # Configurações de infraestrutura
├── scripts/                # Scripts úteis
└── services/               # Microserviços
    ├── catalog/            # Catálogo de produtos
    ├── notification/       # Notificações
    ├── order/              # Gestão de pedidos
    ├── payment/            # Processamento de pagamentos
    └── user/               # Gestão de usuários
```

## 🔧 Comandos Make Úteis

```bash
# Ver todos os comandos disponíveis
make help

# CI/CD local
make ci-local

# Abrir Prometheus
make prometheus

# Abrir Grafana
make grafana

# Abrir Jaeger
make jaeger

# Criar release
make release VERSION=v1.0.0
```

## 🎓 Próximos Passos

1. **Explore a Documentação**
   - [Passo 10: Monitoring e Observability](./PASSO-10-MONITORING.md)
   - [Passo 11: CI/CD e Testes](./PASSO-11-CICD.md)

2. **Customize a Aplicação**
   - Adicione novos endpoints no GraphQL
   - Crie novas queries no Prometheus
   - Configure alertas no Grafana

3. **Deploy em Produção**
   - Configure secrets no GitHub
   - Ajuste limites de recursos no Docker Compose
   - Configure backup do banco de dados

## 🆘 Suporte

- 📖 **Documentação completa**: Ver pasta `docs/`
- 🐛 **Issues**: Abra uma issue no GitHub
- 💬 **Discussões**: Use GitHub Discussions

## 📝 Checklist Inicial

- [ ] Docker e Docker Compose instalados
- [ ] Repositório clonado
- [ ] `docker-compose up -d` executado com sucesso
- [ ] Health check passou (verde)
- [ ] Frontend acessível em http://localhost:3001
- [ ] GraphQL Playground acessível em http://localhost:8080/graphql
- [ ] Grafana acessível em http://localhost:3000
- [ ] Testes executados com sucesso (`make test`)
- [ ] Explorou os dashboards do Grafana
- [ ] Criou uma query de teste no GraphQL

**Parabéns! 🎉 Sua arquitetura de microserviços está rodando!**
