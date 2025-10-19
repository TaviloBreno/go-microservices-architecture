# 🚀 BFF GraphQL - Passo 08

Backend For Frontend (BFF) usando GraphQL que unifica o acesso a todos os microserviços em uma única API.

## 📋 Descrição

Este BFF GraphQL atua como um gateway unificado que:
- 🔗 Conecta com todos os microserviços via gRPC
- 📊 Fornece uma API GraphQL única para o frontend React
- 🎯 Consolida dados de múltiplos serviços em uma única query
- 🔄 Suporta mutations para criação de pedidos e usuários

## 🏗️ Arquitetura

```
Frontend (React)
       ↓ GraphQL HTTP
   BFF GraphQL
       ↓ gRPC
┌─────────────────────────┐
│  Order Service (50052)  │
│  User Service (50051)   │  
│  Payment Service (50053)│
│  Notification (50055)   │
└─────────────────────────┘
```

## 📊 Schema GraphQL

### Tipos Principais

```graphql
type Order {
  id: ID!
  userID: Int!
  productName: String!
  quantity: Int!
  price: Float!
  status: String!
  createdAt: String!
}

type User {
  id: ID!
  name: String!
  email: String!
  createdAt: String!
}

type Payment {
  id: ID!
  orderID: Int!
  userID: Int!
  amount: Float!
  status: String!
  paymentMethod: String!
  createdAt: String!
}

type Notification {
  id: ID!
  orderID: Int!
  userID: Int!
  message: String!
  type: String!
  status: String!
  createdAt: String!
}

type OrderSummary {
  order: Order
  user: User
  payment: Payment
  notifications: [Notification!]!
}
```

### Queries

```graphql
# Listar todos os pedidos
query {
  orders {
    id
    userID
    productName
    quantity
    price
    status
  }
}

# Buscar pedido específico
query {
  order(id: "1") {
    id
    productName
    quantity
    price
  }
}

# Resumo completo do pedido
query {
  orderSummary(orderID: "1") {
    order { id, productName }
    user { name, email }
    payment { status, amount }
    notifications { message, type }
  }
}
```

### Mutations

```graphql
# Criar novo pedido
mutation {
  createOrder(input: {
    userID: 1
    productName: "Produto Exemplo"
    quantity: 2
    price: 99.90
  }) {
    id
    status
  }
}

# Criar novo usuário
mutation {
  createUser(input: {
    name: "João Silva"
    email: "joao@example.com"
  }) {
    id
    name
    email
  }
}
```

## 🛠️ Tecnologias

- **Go 1.21+**: Linguagem principal
- **github.com/99designs/gqlgen**: Geração de código GraphQL
- **gRPC**: Comunicação com microserviços
- **Gorilla WebSocket**: Suporte a subscriptions GraphQL
- **CORS**: Middleware para frontend React
- **Docker**: Containerização

## 🚀 Executando

### Via Docker Compose

```bash
# Subir todos os serviços
docker-compose up -d

# Verificar logs do BFF GraphQL
docker-compose logs -f bff-graphql
```

### Desenvolvimento Local

```bash
# Baixar dependências
go mod tidy

# Executar servidor GraphQL
go run cmd/main.go
```

## 🌐 Endpoints

- **GraphQL API**: `http://localhost:8080/`
- **GraphQL Playground**: `http://localhost:8080/playground`
- **Health Check**: `http://localhost:8080/health`

## ⚙️ Configuração

Variáveis de ambiente disponíveis:

```env
# Porta do servidor GraphQL
PORT=8080

# URLs dos microserviços gRPC
ORDER_SERVICE_URL=localhost:50052
USER_SERVICE_URL=localhost:50051
PAYMENT_SERVICE_URL=localhost:50053
NOTIFICATION_SERVICE_URL=localhost:50055

# Ambiente de execução
GO_ENV=development
```

## 🔧 Desenvolvimento

### Regenerar Código GraphQL

```bash
# Após modificar schema.graphqls
go run github.com/99designs/gqlgen generate
```

### Estrutura de Diretórios

```
bff-graphql/
├── cmd/
│   └── main.go              # Servidor HTTP principal
├── graph/
│   ├── schema.graphqls      # Schema GraphQL
│   ├── schema.resolvers.go  # Resolvers gerados
│   └── resolvers.go         # Implementação dos resolvers
├── internal/
│   ├── clients/             # Clientes gRPC
│   │   ├── order.go
│   │   ├── user.go
│   │   ├── payment.go
│   │   └── notification.go
│   └── config/
│       └── config.go        # Configurações
├── gqlgen.yml              # Configuração do gqlgen
├── go.mod
└── Dockerfile
```

## 📊 Monitoramento

### Health Check

```bash
curl http://localhost:8080/health
```

### Logs de Acesso

O BFF registra todas as requisições com:
- Método HTTP
- Path da URL
- Status Code
- Tempo de resposta
- User-Agent

### Métricas GraphQL

- Cache de queries com LRU
- Suporte a Automatic Persisted Queries (APQ)
- Introspection habilitada em desenvolvimento

## 🔄 Integração com Microserviços

### Conexões gRPC

O BFF mantém conexões persistentes com todos os microserviços:

1. **Order Service** (50052): Gestão de pedidos
2. **User Service** (50051): Gestão de usuários  
3. **Payment Service** (50053): Processamento de pagamentos
4. **Notification Service** (50055): Envio de notificações

### Tratamento de Erros

- Timeout configurável para chamadas gRPC
- Fallback gracioso quando serviços estão indisponíveis
- Logs detalhados para debugging

## 🚦 Status de Desenvolvimento

- ✅ Estrutura do projeto
- ✅ Schema GraphQL completo
- ✅ Clientes gRPC para todos os serviços
- ✅ Resolvers implementados
- ✅ Servidor HTTP com middleware
- ✅ Dockerfile e Docker Compose
- ✅ Documentação completa

## 🎯 Próximos Passos

1. **Frontend React**: Integrar com GraphQL Client (Apollo/Relay)
2. **Autenticação**: JWT tokens e middleware de auth
3. **Subscriptions**: WebSocket para updates em tempo real
4. **Cache**: Implementar cache Redis para performance
5. **Testes**: Testes unitários e de integração

## 🐛 Troubleshooting

### Erro de Conexão gRPC

```bash
# Verificar se os microserviços estão rodando
docker-compose ps

# Testar conectividade
telnet localhost 50051
```

### Playground não Carrega

- Verifique se `GO_ENV=development`
- Acesse `http://localhost:8080/playground`
- Verifique logs do container

### Timeout de Queries

- Ajuste `GRPC_TIMEOUT` (padrão: 10s)
- Verifique performance dos microserviços
- Monitore logs dos resolvers