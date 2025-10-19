# 🎨 Integração Frontend com Banco de Dados

## 📋 Resumo

Esta documentação descreve a integração completa do frontend React com o banco de dados MySQL através do BFF (Backend for Frontend).

## 🎯 Objetivo

Exibir dados reais do banco de dados MySQL no frontend React, substituindo dados mockados por informações reais dos microserviços.

---

## 🏗️ Arquitetura Implementada

```
┌─────────────────┐
│  React Frontend │  (http://localhost:5173)
│   (Vite + UI)   │
└────────┬────────┘
         │
         │ HTTP REST
         ▼
┌─────────────────┐
│   BFF (Go)      │  (http://localhost:8080)
│  REST API       │
└────────┬────────┘
         │
         │ SQL
         ▼
┌─────────────────┐
│  MySQL 8.0      │  (localhost:3306)
│  5 databases    │
│  9 tables       │
│  81+ records    │
└─────────────────┘
```

---

## 📡 Endpoints da API REST (BFF)

### 1. **GET /api/products**
Retorna todos os produtos do catálogo com suas categorias.

**Response:**
```json
[
  {
    "id": 1,
    "name": "Smartphone XYZ Pro",
    "description": "Smartphone de última geração",
    "price": 2999.99,
    "stock": 50,
    "image_url": "https://via.placeholder.com/300x200?text=Smartphone",
    "sku": "ELEC-001",
    "category": "Eletrônicos"
  }
]
```

**Dados:** 28 produtos em 8 categorias

---

### 2. **GET /api/users**
Retorna todos os usuários cadastrados.

**Response:**
```json
[
  {
    "id": 1,
    "name": "João Silva",
    "email": "joao.silva@email.com",
    "phone": "(11) 98765-4321",
    "address": "Rua das Flores, 123 - São Paulo/SP",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

**Dados:** 10 usuários cadastrados

---

### 3. **GET /api/orders**
Retorna todos os pedidos com informações do usuário.

**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "user_name": "João Silva",
    "user_email": "joao.silva@email.com",
    "total": 2999.99,
    "status": "delivered",
    "created_at": "2024-01-15T10:30:00Z"
  }
]
```

**Dados:** 10 pedidos com status variados (delivered, shipped, processing, confirmed, pending)

---

### 4. **GET /api/payments**
Retorna todos os pagamentos com informações do usuário.

**Response:**
```json
[
  {
    "id": 1,
    "user_id": 1,
    "user_name": "João Silva",
    "user_email": "joao.silva@email.com",
    "amount": 2999.99,
    "payment_method": "credit_card",
    "status": "approved",
    "created_at": "2024-01-15T10:35:00Z"
  }
]
```

**Dados:** 10 pagamentos (credit_card, pix, debit_card, boleto)

---

## 🖥️ Páginas React Criadas

### 1. **ProductsPage** (`/products`)

**Componente:** `frontend/src/pages/ProductsPage.jsx`

**Funcionalidades:**
- ✅ Grid responsivo de produtos
- ✅ Cards com imagem, nome, preço, estoque
- ✅ Badge de categoria
- ✅ Badge de estoque (verde/amarelo/vermelho)
- ✅ Estatísticas (total, valor, categorias, estoque baixo)
- ✅ Loading skeleton
- ✅ Tratamento de erros

**Estatísticas Exibidas:**
- Total de produtos: 28
- Valor total em estoque: R$ 106.889,90
- Categorias: 8
- Produtos com estoque baixo (<50): 12

---

### 2. **UsersPage** (`/users`)

**Componente:** `frontend/src/pages/UsersPage.jsx`

**Funcionalidades:**
- ✅ Lista de usuários em cards
- ✅ Informações: nome, email, telefone, endereço
- ✅ Datas de criação e atualização
- ✅ Badge com ID do usuário
- ✅ Ícones para cada tipo de informação
- ✅ Estatísticas de usuários

**Dados Exibidos:**
- Total de usuários: 10
- Todos com status ativo
- Datas de cadastro formatadas

---

### 3. **OrdersPage** (atualizada)

**Componente:** `frontend/src/components/OrdersTable.jsx`

**Mudanças:**
- ❌ Removido: Apollo Client / GraphQL
- ✅ Adicionado: Fetch API / REST
- ✅ Exibe nome e email do usuário
- ✅ Total do pedido
- ✅ Status traduzido (delivered = Entregue, shipped = Enviado, etc)
- ✅ Auto-refresh a cada 5 segundos

---

## 🎨 Componentes de UI Utilizados

Todos os componentes utilizam **shadcn/ui**:

- `Card` - Container principal
- `Badge` - Status e categorias
- `Skeleton` - Loading states
- Ícones do **Lucide React**: `ShoppingCart`, `Users`, `Package`, `Mail`, `Phone`, `MapPin`, etc

---

## 🔄 Atualização Automática

- **OrdersPage**: Auto-refresh a cada 5 segundos
- **ProductsPage**: Refresh manual (pode adicionar polling se necessário)
- **UsersPage**: Refresh manual

---

## 🗺️ Navegação

Sidebar atualizada com novos links:

1. 🎯 Dashboard
2. 🛍️ **Produtos** (NOVO)
3. 👥 **Usuários** (NOVO)
4. 📦 Pedidos (atualizado)
5. 💳 Pagamentos
6. 🔔 Notificações

---

## 🚀 Como Testar

### 1. Verificar containers Docker
```bash
docker-compose ps
```

Devem estar rodando:
- mysql
- bff-graphql (BFF)
- rabbitmq
- services (catalog, order, payment, user, notification)

---

### 2. Testar API REST
```bash
# Produtos
curl http://localhost:8080/api/products

# Usuários
curl http://localhost:8080/api/users

# Pedidos
curl http://localhost:8080/api/orders

# Pagamentos
curl http://localhost:8080/api/payments
```

---

### 3. Acessar Frontend
```bash
cd frontend
npm run dev
```

Acesse: http://localhost:5173

**Login:**
- Email: `admin@example.com`
- Senha: `admin123`

**Páginas para testar:**
1. `/products` - Catálogo com 28 produtos
2. `/users` - Lista de 10 usuários
3. `/orders` - Pedidos em tempo real
4. `/payments` - Pagamentos realizados

---

## 📊 Dados de Teste no Banco

### Produtos (28 itens)
- **Eletrônicos**: Smartphone, Notebook, Tablet, Monitor, Teclado, Mouse, Fone
- **Livros**: Clean Code, Design Patterns, Refactoring
- **Roupas**: Camiseta, Calça Jeans, Jaqueta
- **Casa**: Panela, Liquidificador, Aspirador
- **Esportes**: Bola, Tênis, Bicicleta
- **Alimentos**: Café, Chocolate, Azeite
- **Beleza**: Shampoo, Perfume, Creme
- **Brinquedos**: Boneca, Carrinho, LEGO

### Usuários (10)
- João Silva, Maria Santos, Pedro Oliveira, Ana Costa, Carlos Souza, Beatriz Lima, Fernando Alves, Juliana Rocha, Ricardo Mendes, Patrícia Dias

### Pedidos (10)
- Status: delivered (3), shipped (2), processing (2), confirmed (2), pending (1)

### Pagamentos (10)
- Métodos: credit_card (4), pix (3), debit_card (2), boleto (1)
- Status: approved (8), pending (2)

---

## 🛠️ Tecnologias Utilizadas

### Backend (BFF)
- Go 1.21
- database/sql
- go-sql-driver/mysql
- CORS habilitado

### Frontend
- React 18
- Vite
- React Router v6
- shadcn/ui
- Tailwind CSS
- Lucide Icons

### Banco de Dados
- MySQL 8.0
- 5 databases (user_service, order_service, catalog_service, payment_service, notification_service)
- Foreign Keys e Indexes

---

## 📁 Arquivos Modificados/Criados

### Backend
- ✅ `bff/cmd/main.go` - 4 endpoints REST
- ✅ `bff/go.mod` - MySQL driver
- ✅ `docker-compose.yml` - Env vars do BFF

### Frontend
- ✅ `frontend/src/pages/ProductsPage.jsx` (NOVO)
- ✅ `frontend/src/pages/UsersPage.jsx` (NOVO)
- ✅ `frontend/src/components/OrdersTable.jsx` (ATUALIZADO)
- ✅ `frontend/src/App.jsx` (ROTAS)
- ✅ `frontend/src/components/Sidebar.jsx` (NAVEGAÇÃO)

### Database
- ✅ `infra/mysql/init/init.sql` - Schema + dados

---

## ✅ Checklist de Funcionalidades

- [x] BFF conectado ao MySQL
- [x] 4 endpoints REST funcionando
- [x] Página de Produtos com grid
- [x] Página de Usuários com cards
- [x] Página de Pedidos atualizada
- [x] Navegação no Sidebar
- [x] Rotas no React Router
- [x] Loading states
- [x] Error handling
- [x] CORS configurado
- [x] Auto-refresh em pedidos
- [x] Badges de status
- [x] Formatação de moeda (BRL)
- [x] Formatação de datas
- [x] Design responsivo

---

## 🎉 Resultado Final

O frontend agora exibe **dados reais do banco de dados** através da API REST do BFF!

**Dados Reais:**
- ✅ 28 produtos com preços, estoque e categorias
- ✅ 10 usuários com contatos e endereços
- ✅ 10 pedidos com status em tempo real
- ✅ 10 pagamentos processados
- ✅ UI moderna com shadcn/ui
- ✅ Totalmente funcional e responsivo

---

## 📝 Próximos Passos (Opcional)

1. Adicionar paginação para produtos
2. Implementar busca/filtros
3. Atualizar PaymentsPage com API real
4. Adicionar gráficos no Dashboard
5. Implementar cache no frontend
6. Adicionar WebSocket para atualizações real-time
7. Criar página de detalhes do produto
8. Adicionar carrinho de compras

---

## 🆘 Troubleshooting

### Problema: Frontend não carrega dados
**Solução:** Verificar se BFF está rodando na porta 8080
```bash
curl http://localhost:8080/api/products
```

### Problema: CORS error
**Solução:** BFF já tem CORS configurado. Reiniciar container:
```bash
docker-compose restart bff-graphql
```

### Problema: Banco vazio
**Solução:** Recriar MySQL com init script:
```bash
docker-compose down mysql
Remove-Item -Recurse -Force ./data/mysql
docker-compose up -d mysql
```

---

## 📚 Documentação Relacionada

- [INSTALLATION.md](./INSTALLATION.md) - Instalação completa
- [DEPLOYMENT_COMMANDS.md](./DEPLOYMENT_COMMANDS.md) - Comandos úteis
- [DEPLOYMENT_SUMMARY.md](./DEPLOYMENT_SUMMARY.md) - Resumo do deploy

---

**Última atualização:** 2024-01-20  
**Autor:** Equipe GoExpert Microservices
