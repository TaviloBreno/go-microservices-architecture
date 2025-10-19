# 🏛️ Arquitetura de Microserviços - Documentação Técnica

## 📋 Índice

1. [Visão Geral](#-visão-geral)
2. [Arquitetura do Sistema](#-arquitetura-do-sistema)
3. [Microserviços](#-microserviços)
4. [Comunicação Entre Serviços](#-comunicação-entre-serviços)
5. [Banco de Dados](#-banco-de-dados)
6. [Message Queue](#-message-queue)
7. [Monitoramento e Observabilidade](#-monitoramento-e-observabilidade)
8. [Segurança](#-segurança)
9. [Escalabilidade](#-escalabilidade)
10. [Padrões e Práticas](#-padrões-e-práticas)

---

## 🎯 Visão Geral

Esta é uma arquitetura completa de microserviços em Go, demonstrando as melhores práticas de desenvolvimento, comunicação, monitoramento e deployment.

### Características Principais

- **Microserviços** independentes e escaláveis
- **gRPC** para comunicação síncrona de alta performance
- **GraphQL BFF** para agregação de dados e API unificada
- **Message Queue** (RabbitMQ) para comunicação assíncrona
- **Observabilidade completa** com Prometheus, Grafana e Jaeger
- **CI/CD** automatizado com GitHub Actions
- **Testes automatizados** com cobertura de código
- **Containerização** completa com Docker

### Stack Tecnológica

| Categoria | Tecnologia | Versão | Finalidade |
|-----------|-----------|--------|------------|
| **Backend** | Go | 1.21+ | Linguagem principal |
| **Frontend** | React | 18+ | Interface do usuário |
| **API Gateway** | GraphQL | - | BFF (Backend for Frontend) |
| **RPC** | gRPC | - | Comunicação entre microserviços |
| **Database** | MySQL | 8.0+ | Persistência de dados |
| **Message Broker** | RabbitMQ | 3.12+ | Comunicação assíncrona |
| **Metrics** | Prometheus | 2.48+ | Coleta de métricas |
| **Visualization** | Grafana | 10.2+ | Dashboards e visualização |
| **Tracing** | Jaeger | 1.55+ | Rastreamento distribuído |
| **Containerization** | Docker | 20.10+ | Containerização |
| **Orchestration** | Docker Compose | 2.0+ | Orquestração local |
| **CI/CD** | GitHub Actions | - | Integração e deployment contínuos |

---

## 🏗️ Arquitetura do Sistema

### Diagrama de Alto Nível

```
┌─────────────────────────────────────────────────────────────────┐
│                          FRONTEND LAYER                          │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │         React Dashboard (Port 3001)                      │   │
│  │  - Dark Mode Support                                     │   │
│  │  - Real-time Updates                                     │   │
│  │  - GraphQL Client                                        │   │
│  └───────────────────────────┬─────────────────────────────┘   │
└────────────────────────────────┼─────────────────────────────────┘
                                 │ HTTP/GraphQL
                                 ▼
┌─────────────────────────────────────────────────────────────────┐
│                         BFF LAYER (API Gateway)                  │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │    GraphQL BFF (Port 8080)                              │   │
│  │  - Query/Mutation Resolvers                             │   │
│  │  - Service Aggregation                                  │   │
│  │  - Authentication/Authorization                         │   │
│  └────┬────────┬─────────┬──────────┬────────┬────────────┘   │
└───────┼────────┼─────────┼──────────┼────────┼──────────────────┘
        │        │         │          │        │
        │ gRPC   │ gRPC    │ gRPC     │ gRPC   │ gRPC
        ▼        ▼         ▼          ▼        ▼
┌─────────────────────────────────────────────────────────────────┐
│                      MICROSERVICES LAYER                         │
│                                                                   │
│  ┌───────────┐  ┌───────────┐  ┌───────────┐  ┌──────────────┐│
│  │  Order    │  │  Payment  │  │   User    │  │ Notification ││
│  │  Service  │  │  Service  │  │  Service  │  │   Service    ││
│  │ :50051    │  │ :50052    │  │ :50053    │  │   :50054     ││
│  └─────┬─────┘  └─────┬─────┘  └─────┬─────┘  └──────┬───────┘│
│        │              │              │                │         │
│        │              │              │                │         │
│  ┌───────────┐        │              │                │         │
│  │  Catalog  │        │              │                │         │
│  │  Service  │        │              │                │         │
│  │  :50055   │        │              │                │         │
│  └─────┬─────┘        │              │                │         │
│        │              │              │                │         │
└────────┼──────────────┼──────────────┼────────────────┼─────────┘
         │              │              │                │
         ├──────────────┴──────────────┴────────────────┘
         │                     MySQL                     
         │                   (Port 3306)                 
         │                                               
         └────────────────────────────────────────────────┐
                                                           │
┌──────────────────────────────────────────────────────────┼──────┐
│                     MESSAGING LAYER                      │       │
│                                                           ▼       │
│  ┌────────────────────────────────────────────────────────────┐ │
│  │              RabbitMQ (Port 5672)                          │ │
│  │  - Order Events Queue                                      │ │
│  │  - Payment Events Queue                                    │ │
│  │  - Notification Queue                                      │ │
│  └────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                   OBSERVABILITY LAYER                            │
│                                                                   │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │  Prometheus  │  │   Grafana    │  │    Jaeger    │          │
│  │   :9090      │  │    :3000     │  │   :16686     │          │
│  │              │  │              │  │              │          │
│  │  Metrics     │  │  Dashboards  │  │  Tracing     │          │
│  │  Collection  │  │  Alerts      │  │  Analysis    │          │
│  └──────────────┘  └──────────────┘  └──────────────┘          │
└─────────────────────────────────────────────────────────────────┘
```

### Fluxo de Dados Principal

1. **Cliente → Frontend**: Usuário interage com React Dashboard
2. **Frontend → BFF**: Requisição GraphQL via HTTP
3. **BFF → Microserviços**: Chamadas gRPC paralelas para múltiplos serviços
4. **Microserviços → Database**: Operações CRUD no MySQL
5. **Microserviços → Message Queue**: Publicação de eventos assíncronos
6. **Consumers → Notification**: Processamento de eventos e notificações
7. **Todos → Observability**: Métricas, logs e traces para monitoramento

---

## 🔧 Microserviços

### 1. Order Service (Pedidos) - Port 50051

**Responsabilidades:**
- Criação e gestão de pedidos
- Validação de itens do pedido
- Cálculo de totais
- Publicação de eventos de pedido

**Endpoints gRPC:**
```protobuf
service OrderService {
  rpc CreateOrder(CreateOrderRequest) returns (Order);
  rpc GetOrder(GetOrderRequest) returns (Order);
  rpc ListOrders(ListOrdersRequest) returns (OrderList);
  rpc UpdateOrderStatus(UpdateOrderStatusRequest) returns (Order);
  rpc CancelOrder(CancelOrderRequest) returns (Order);
}
```

**Eventos Publicados:**
- `order.created` - Quando um pedido é criado
- `order.updated` - Quando status do pedido muda
- `order.cancelled` - Quando pedido é cancelado

**Métricas Expostas:**
- `orders_created_total` - Total de pedidos criados
- `orders_by_status` - Pedidos agrupados por status
- `order_value_total` - Valor total de pedidos

**Tecnologias:**
- gRPC Server
- MySQL (tabela: `orders`, `order_items`)
- RabbitMQ Publisher
- Prometheus Client
- OpenTelemetry

---

### 2. Payment Service (Pagamentos) - Port 50052

**Responsabilidades:**
- Processamento de pagamentos
- Validação de métodos de pagamento
- Cálculo de taxas
- Integração com gateways de pagamento (simulado)

**Endpoints gRPC:**
```protobuf
service PaymentService {
  rpc ProcessPayment(ProcessPaymentRequest) returns (Payment);
  rpc GetPayment(GetPaymentRequest) returns (Payment);
  rpc RefundPayment(RefundPaymentRequest) returns (Payment);
  rpc ValidatePaymentMethod(ValidateRequest) returns (ValidationResult);
}
```

**Métodos de Pagamento Suportados:**
- `credit_card` - Cartão de crédito
- `debit_card` - Cartão de débito
- `pix` - PIX
- `boleto` - Boleto bancário

**Eventos Publicados:**
- `payment.processed` - Pagamento processado com sucesso
- `payment.failed` - Falha no processamento
- `payment.refunded` - Pagamento reembolsado

**Métricas Expostas:**
- `payments_processed_total` - Total de pagamentos processados
- `payments_by_method` - Pagamentos por método
- `payment_amount_total` - Valor total processado
- `payment_fees_total` - Total de taxas cobradas

---

### 3. User Service (Usuários) - Port 50053

**Responsabilidades:**
- Cadastro e autenticação de usuários
- Gestão de perfis
- Validação de dados
- Controle de acesso

**Endpoints gRPC:**
```protobuf
service UserService {
  rpc CreateUser(CreateUserRequest) returns (User);
  rpc GetUser(GetUserRequest) returns (User);
  rpc UpdateUser(UpdateUserRequest) returns (User);
  rpc DeleteUser(DeleteUserRequest) returns (DeleteResponse);
  rpc ListUsers(ListUsersRequest) returns (UserList);
  rpc AuthenticateUser(AuthRequest) returns (AuthResponse);
}
```

**Validações:**
- Email único e formato válido
- Senha forte (mínimo 8 caracteres)
- Campos obrigatórios

**Eventos Publicados:**
- `user.created` - Novo usuário criado
- `user.updated` - Dados do usuário atualizados
- `user.deleted` - Usuário removido

**Métricas Expostas:**
- `users_created_total` - Total de usuários criados
- `users_active_total` - Usuários ativos
- `user_logins_total` - Total de logins

---

### 4. Notification Service (Notificações) - Port 50054

**Responsabilidades:**
- Envio de notificações por email, SMS, push
- Processamento assíncrono de eventos
- Templates de mensagens
- Histórico de notificações

**Endpoints gRPC:**
```protobuf
service NotificationService {
  rpc SendNotification(SendNotificationRequest) returns (NotificationResponse);
  rpc GetNotification(GetNotificationRequest) returns (Notification);
  rpc ListNotifications(ListNotificationsRequest) returns (NotificationList);
}
```

**Tipos de Notificação:**
- `email` - Email transacional
- `sms` - Mensagem de texto
- `push` - Notificação push
- `webhook` - Webhook HTTP

**Eventos Consumidos:**
- `order.created` → Confirma pedido por email
- `payment.processed` → Confirma pagamento
- `user.created` → Email de boas-vindas

**Métricas Expostas:**
- `notifications_sent_total` - Notificações enviadas
- `notifications_by_type` - Por tipo de notificação
- `notification_failures_total` - Falhas no envio

---

### 5. Catalog Service (Catálogo) - Port 50055

**Responsabilidades:**
- Gestão de produtos
- Controle de estoque
- Categorização
- Busca e filtros

**Endpoints gRPC:**
```protobuf
service CatalogService {
  rpc CreateProduct(CreateProductRequest) returns (Product);
  rpc GetProduct(GetProductRequest) returns (Product);
  rpc ListProducts(ListProductsRequest) returns (ProductList);
  rpc UpdateStock(UpdateStockRequest) returns (Product);
  rpc SearchProducts(SearchRequest) returns (ProductList);
}
```

**Funcionalidades:**
- CRUD completo de produtos
- Gestão de categorias
- Controle de estoque em tempo real
- Busca por nome, categoria, preço

**Eventos Publicados:**
- `product.created` - Produto criado
- `product.updated` - Produto atualizado
- `stock.updated` - Estoque alterado

**Métricas Expostas:**
- `products_total` - Total de produtos
- `products_out_of_stock` - Produtos sem estoque
- `catalog_searches_total` - Total de buscas

---

### 6. BFF GraphQL (Backend for Frontend) - Port 8080

**Responsabilidades:**
- Agregação de dados de múltiplos microserviços
- API unificada para o frontend
- Otimização de queries
- Cache de respostas

**Schema GraphQL Principal:**

```graphql
type Query {
  # Users
  user(id: ID!): User
  users: [User!]!
  
  # Orders
  order(id: ID!): Order
  orders(userId: ID): [Order!]!
  
  # Payments
  payment(id: ID!): Payment
  payments(orderId: ID): [Payment!]!
  
  # Products
  product(id: ID!): Product
  products(category: String): [Product!]!
  
  # Notifications
  notifications(userId: ID!): [Notification!]!
}

type Mutation {
  # Users
  createUser(input: CreateUserInput!): User!
  updateUser(id: ID!, input: UpdateUserInput!): User!
  
  # Orders
  createOrder(input: CreateOrderInput!): Order!
  cancelOrder(id: ID!): Order!
  
  # Payments
  processPayment(input: ProcessPaymentInput!): Payment!
  refundPayment(id: ID!): Payment!
  
  # Products
  createProduct(input: CreateProductInput!): Product!
  updateStock(id: ID!, quantity: Int!): Product!
}
```

**Otimizações:**
- DataLoader para batch requests
- Query complexity analysis
- Response caching
- Rate limiting

---

## 📡 Comunicação Entre Serviços

### Síncrona (gRPC)

**Por que gRPC?**
- ✅ Alta performance (Protocol Buffers)
- ✅ Type-safe (contratos .proto)
- ✅ Streaming bidirecional
- ✅ Suporte a múltiplas linguagens
- ✅ HTTP/2

**Exemplo de Comunicação:**

```go
// BFF chamando Order Service
conn, _ := grpc.Dial("order-service:50051", grpc.WithInsecure())
client := orderpb.NewOrderServiceClient(conn)

ctx := context.Background()
order, err := client.CreateOrder(ctx, &orderpb.CreateOrderRequest{
    UserId: "123",
    Items: []*orderpb.OrderItem{
        {ProductId: "p1", Quantity: 2, Price: 99.90},
    },
    TotalAmount: 199.80,
})
```

**Interceptors Implementados:**
- **Logging**: Log de todas as chamadas
- **Metrics**: Contadores e histogramas
- **Tracing**: Propagação de context com OpenTelemetry
- **Error Handling**: Tratamento padronizado de erros
- **Authentication**: Validação de tokens JWT

---

### Assíncrona (RabbitMQ)

**Por que Message Queue?**
- ✅ Desacoplamento de serviços
- ✅ Processamento assíncrono
- ✅ Tolerância a falhas
- ✅ Balanceamento de carga
- ✅ Garantia de entrega

**Topologia:**

```
┌─────────────┐         ┌──────────────┐         ┌─────────────┐
│   Order     │         │   Exchange   │         │ Notification│
│   Service   ├────────►│   (topic)    ├────────►│   Service   │
└─────────────┘         │              │         └─────────────┘
    Publish             │  Routing:    │             Subscribe
                        │  order.*     │
                        │  payment.*   │
┌─────────────┐         │  user.*      │         ┌─────────────┐
│   Payment   │         │              │         │   Audit     │
│   Service   ├────────►│              ├────────►│   Service   │
└─────────────┘         └──────────────┘         └─────────────┘
```

**Padrões Implementados:**
- **Publisher/Subscriber**: Eventos de domínio
- **Work Queue**: Processamento de tarefas
- **Dead Letter Queue**: Mensagens com falha
- **Retry Logic**: Tentativas automáticas

**Exemplo de Publicação:**

```go
// Publicar evento de pedido criado
event := OrderCreatedEvent{
    OrderID:     order.ID,
    UserID:      order.UserID,
    TotalAmount: order.TotalAmount,
    CreatedAt:   time.Now(),
}

data, _ := json.Marshal(event)
err := channel.Publish(
    "events",           // exchange
    "order.created",    // routing key
    false,              // mandatory
    false,              // immediate
    amqp.Publishing{
        ContentType: "application/json",
        Body:        data,
    },
)
```

---

## 🗄️ Banco de Dados

### MySQL Schema

**Principais Tabelas:**

```sql
-- Usuários
CREATE TABLE users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email)
);

-- Produtos
CREATE TABLE products (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INT NOT NULL DEFAULT 0,
    category VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_category (category),
    INDEX idx_price (price)
);

-- Pedidos
CREATE TABLE orders (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    total_amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_user_id (user_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
);

-- Itens do Pedido
CREATE TABLE order_items (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    product_id VARCHAR(36) NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id),
    INDEX idx_order_id (order_id)
);

-- Pagamentos
CREATE TABLE payments (
    id VARCHAR(36) PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    transaction_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    processed_at TIMESTAMP NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    INDEX idx_order_id (order_id),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
);

-- Notificações
CREATE TABLE notifications (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(36) NOT NULL,
    type VARCHAR(50) NOT NULL,
    subject VARCHAR(255),
    message TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    sent_at TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    INDEX idx_user_id (user_id),
    INDEX idx_status (status),
    INDEX idx_type (type)
);
```

**Padrões Aplicados:**
- ✅ UUID como chave primária
- ✅ Timestamps automáticos
- ✅ Índices em foreign keys
- ✅ Constraints de integridade referencial
- ✅ Soft deletes onde aplicável

---

## 📊 Monitoramento e Observabilidade

### Prometheus (Métricas)

**Métricas Coletadas:**

1. **Sistema:**
   - CPU usage
   - Memory usage
   - Disk I/O
   - Network traffic

2. **Aplicação:**
   - Request rate
   - Error rate
   - Response time (latência)
   - Throughput

3. **gRPC:**
   - `grpc_server_started_total` - Total de requests iniciadas
   - `grpc_server_handled_total` - Total de requests completadas
   - `grpc_server_handling_seconds` - Latência das requests

4. **Business:**
   - `orders_created_total`
   - `payments_processed_total`
   - `users_active_total`
   - `notifications_sent_total`

**Queries Úteis:**

```promql
# Taxa de erro por serviço
rate(grpc_server_handled_total{grpc_code!="OK"}[5m])

# Latência P95
histogram_quantile(0.95, rate(grpc_server_handling_seconds_bucket[5m]))

# Throughput por método
rate(grpc_server_handled_total[5m])

# Taxa de sucesso
sum(rate(grpc_server_handled_total{grpc_code="OK"}[5m])) 
/ 
sum(rate(grpc_server_handled_total[5m]))
```

---

### Grafana (Dashboards)

**Dashboards Criados:**

1. **Microservices Overview**
   - Status de todos os serviços
   - Request rate geral
   - Error rate geral
   - CPU e memória por serviço

2. **gRPC Performance**
   - Latência por método
   - Taxa de sucesso/erro
   - Throughput
   - Conexões ativas

3. **Business Metrics**
   - Pedidos por status
   - Receita total
   - Conversão de pagamentos
   - Usuários ativos

**Alertas Configurados:**
- High Error Rate (> 5%)
- High Latency (P95 > 1s)
- Service Down
- Database Connection Issues

---

### Jaeger (Distributed Tracing)

**Spans Rastreados:**

```
GET /graphql?query=users
├─ GraphQL: users Query                [BFF]
   ├─ gRPC: UserService.ListUsers      [User Service]
   │  ├─ MySQL: SELECT FROM users      [Database]
   │  └─ Cache: Check user cache       [Redis]
   └─ Metrics: Record query latency    [Prometheus]
```

**Informações Capturadas:**
- Trace ID (identificador único)
- Span ID (identificador do segmento)
- Parent Span ID (hierarquia)
- Duration (duração)
- Tags (metadata)
- Logs (eventos)
- Baggage (contexto propagado)

---

## 🔒 Segurança

### Autenticação e Autorização

**JWT Tokens:**
```go
type Claims struct {
    UserID string   `json:"user_id"`
    Email  string   `json:"email"`
    Role   string   `json:"role"`
    jwt.StandardClaims
}
```

**Níveis de Acesso:**
- `admin` - Acesso total
- `user` - Operações básicas
- `guest` - Apenas leitura

### Validação de Entrada

```go
// Validação de email
func ValidateEmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}

// Validação de senha forte
func ValidatePassword(password string) error {
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters")
    }
    // Adicionar mais regras...
    return nil
}
```

### Rate Limiting

```go
// Limitar 100 requests por minuto por IP
limiter := rate.NewLimiter(100, 10)
```

### CORS

```go
cors.AllowAll() // Desenvolvimento
// Produção: configurar origens permitidas
```

---

## 📈 Escalabilidade

### Horizontal Scaling

**Stateless Services:**
- Todos os microserviços são stateless
- Escalam horizontalmente com Docker Compose/Kubernetes
- Load balancing automático

**Exemplo com Docker Compose:**
```yaml
services:
  order-service:
    deploy:
      replicas: 3
      resources:
        limits:
          cpus: '0.5'
          memory: 512M
```

### Caching

**Estratégias:**
- Cache de queries frequentes
- TTL configurável
- Invalidação por evento

### Database Optimization

**Índices:**
- Primary keys (UUIDs)
- Foreign keys
- Campos de busca frequente

**Connection Pooling:**
```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

---

## 🎨 Padrões e Práticas

### Design Patterns Implementados

1. **Repository Pattern**
   ```go
   type OrderRepository interface {
       Create(ctx context.Context, order *Order) error
       FindByID(ctx context.Context, id string) (*Order, error)
       FindAll(ctx context.Context) ([]*Order, error)
   }
   ```

2. **Service Pattern**
   ```go
   type OrderService struct {
       repo      OrderRepository
       publisher EventPublisher
       metrics   Metrics
   }
   ```

3. **Factory Pattern**
   ```go
   func NewOrderService(repo OrderRepository) *OrderService {
       return &OrderService{repo: repo}
   }
   ```

4. **Dependency Injection**
   ```go
   func main() {
       repo := NewMySQLRepository(db)
       publisher := NewRabbitMQPublisher(conn)
       service := NewOrderService(repo, publisher)
   }
   ```

### Clean Architecture

```
cmd/                    # Entry point
internal/
  ├─ domain/            # Entities, Value Objects
  ├─ usecase/           # Business Logic
  ├─ repository/        # Data Access
  ├─ handler/           # gRPC Handlers
  └─ infra/             # External dependencies
```

### Error Handling

```go
// Erros customizados
var (
    ErrNotFound = errors.New("resource not found")
    ErrInvalidInput = errors.New("invalid input")
    ErrUnauthorized = errors.New("unauthorized")
)

// Wrapping de erros
if err != nil {
    return fmt.Errorf("failed to create order: %w", err)
}
```

### Testing

**Test Pyramid:**
- 70% Unit Tests
- 20% Integration Tests
- 10% E2E Tests

**Coverage Targets:**
- Minimum: 70%
- Target: 80%+

---

## 📚 Referências

- [Go Best Practices](https://golang.org/doc/effective_go)
- [gRPC Go Documentation](https://grpc.io/docs/languages/go/)
- [GraphQL Best Practices](https://graphql.org/learn/best-practices/)
- [Prometheus Best Practices](https://prometheus.io/docs/practices/)
- [The Twelve-Factor App](https://12factor.net/)
- [Microservices Patterns](https://microservices.io/patterns/)

---

**Última Atualização:** 2024
**Versão:** 1.0.0
**Autor:** Go Expert - Microservices Architecture Team
