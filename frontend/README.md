# 🎨 Frontend React - Passo 09

Interface web moderna construída em React que consome a API GraphQL do BFF para exibir dados dos microserviços em tempo real.

## 📋 Descrição

Este frontend React oferece uma interface intuitiva para:
- 📊 **Dashboard**: Visão geral com estatísticas em tempo real
- 📦 **Pedidos**: Listagem e monitoramento de pedidos
- 💳 **Pagamentos**: Acompanhamento de transações financeiras
- 🔔 **Notificações**: Histórico de notificações enviadas
- 🔄 **Atualizações Automáticas**: Polling a cada 5 segundos

## 🏗️ Tecnologias

- **React 18**: Biblioteca principal com hooks modernos
- **Vite**: Build tool rápido e otimizado
- **Apollo Client**: Client GraphQL com cache inteligente
- **React Router**: Navegação SPA
- **Tailwind CSS**: Framework CSS utility-first
- **PostCSS**: Processamento CSS
- **Nginx**: Servidor web para produção

## 🚀 Como Executar

### 📦 Desenvolvimento Local

```bash
# Instalar dependências
npm install

# Executar servidor de desenvolvimento
npm run dev

# Acessar aplicação
http://localhost:5173
```

### 🐳 Docker (Produção)

```bash
# Build e execução via Docker Compose
docker-compose up -d frontend

# Acessar aplicação
http://localhost:3000
```

## 📊 Estrutura do Projeto

```
frontend/
├── src/
│   ├── apollo/
│   │   ├── client.js           # Configuração Apollo Client
│   │   └── queries.js          # Queries e Mutations GraphQL
│   ├── components/
│   │   ├── Header.jsx          # Cabeçalho com navegação
│   │   ├── OrdersTable.jsx     # Tabela de pedidos
│   │   ├── PaymentsTable.jsx   # Tabela de pagamentos
│   │   └── NotificationsTable.jsx  # Tabela de notificações
│   ├── pages/
│   │   ├── Dashboard.jsx       # Página principal
│   │   ├── OrdersPage.jsx      # Página de pedidos
│   │   ├── PaymentsPage.jsx    # Página de pagamentos
│   │   └── NotificationsPage.jsx   # Página de notificações
│   ├── App.jsx                 # Componente raiz
│   ├── main.jsx               # Entry point
│   └── index.css              # Estilos Tailwind
├── public/                     # Assets estáticos
├── package.json               # Dependências
├── tailwind.config.js         # Configuração Tailwind
├── vite.config.js            # Configuração Vite
└── Dockerfile                # Build para produção
```

## 🔗 Integração GraphQL

### 🌐 Endpoint
```
http://localhost:8080/graphql
```

### 📝 Queries Principais

```graphql
# Dashboard - Todos os dados
query GetDashboardData {
  orders { id userID productName quantity price status createdAt }
  payments { id orderID amount status paymentMethod createdAt }
  notifications { id orderID message type status createdAt }
}

# Resumo de Pedido Completo
query GetOrderSummary($orderID: ID!) {
  orderSummary(orderID: $orderID) {
    order { id productName price status }
    user { name email }
    payment { status amount }
    notifications { message type status }
  }
}
```

### 🔄 Polling e Atualizações

- **Dashboard**: Atualiza a cada 10 segundos
- **Tabelas**: Atualizam a cada 5 segundos
- **Health Check**: Verifica a cada 30 segundos
- **Cache**: Apollo Client com políticas otimizadas

## 🎨 Design System

### 🎯 Cores Principais

```css
/* Primary */
--primary-600: #2563eb;
--primary-100: #dbeafe;

/* Success */
--success-600: #16a34a;
--success-100: #dcfce7;

/* Warning */
--warning-600: #d97706;
--warning-100: #fef3c7;

/* Danger */
--danger-600: #dc2626;
--danger-100: #fee2e2;
```

### 🧩 Componentes

- **Buttons**: `.btn-primary`, `.btn-secondary`
- **Cards**: `.card` com sombra e border-radius
- **Tables**: `.table-container` responsivas
- **Status Badges**: `.status-badge` com cores semânticas
- **Loading**: `.loading-spinner` animado

## 📱 Responsividade

- **Mobile First**: Design adaptativo
- **Breakpoints**: `sm:`, `md:`, `lg:`, `xl:`
- **Grid System**: Flexbox e CSS Grid
- **Touch Friendly**: Botões e links otimizados

## 🔧 Configuração Apollo Client

```javascript
const client = new ApolloClient({
  uri: 'http://localhost:8080/graphql',
  cache: new InMemoryCache({
    typePolicies: {
      Query: {
        fields: {
          orders: { fetchPolicy: 'cache-and-network' },
          payments: { fetchPolicy: 'cache-and-network' },
          notifications: { fetchPolicy: 'cache-and-network' }
        }
      }
    }
  }),
  defaultOptions: {
    watchQuery: {
      pollInterval: 5000,
      errorPolicy: 'all'
    }
  }
});
```

## 🚦 Status e Indicadores

### 🟢 Indicadores de Status

- **Online**: Serviço funcionando
- **Offline**: Serviço indisponível  
- **Conectando**: Tentando conexão
- **Erro**: Falha na comunicação

### 📊 Health Check

- Verifica conectividade com BFF GraphQL
- Mostra status dos microserviços
- Indicador visual na header

## 🔄 Fluxo de Dados

```
Frontend React (Porto 3000)
       ↓ HTTP/GraphQL
   BFF GraphQL (Porto 8080)
       ↓ gRPC
┌─────────────────────────┐
│  Order Service (50052)  │
│  User Service (50051)   │  
│  Payment Service (50053)│
│  Notification (50055)   │
└─────────────────────────┘
```

## 📈 Métricas e Performance

### ⚡ Otimizações

- **Code Splitting**: Lazy loading de páginas
- **Bundle Splitting**: Vendor chunks separados
- **Gzip**: Compressão no Nginx
- **Cache Headers**: Assets com cache longo
- **Tree Shaking**: Remoção de código não usado

### 📊 Performance

- **First Contentful Paint**: < 1.5s
- **Largest Contentful Paint**: < 2.5s  
- **Bundle Size**: < 500KB gzipped
- **Lighthouse Score**: > 90

## 🐛 Tratamento de Erros

### 🔧 Error Boundaries

- Fallback UI para erros React
- Log de erros para debugging
- Retry automático para falhas de rede

### 📡 GraphQL Errors

- Error policy: `all` (mostra dados parciais)
- Retry automático com exponential backoff
- Fallback para dados em cache

## 🧪 Testing

### 🧪 Scripts Disponíveis

```bash
# Desenvolvimento
npm run dev

# Build produção
npm run build

# Preview build
npm run preview

# Lint
npm run lint
```

## 🔒 Segurança

### 🛡️ Headers de Segurança

```nginx
add_header X-Frame-Options DENY;
add_header X-Content-Type-Options nosniff;
add_header X-XSS-Protection "1; mode=block";
```

### 🔐 CORS

- Configurado no BFF GraphQL
- Permite origem do frontend
- Headers necessários incluídos

## 🚀 Deploy

### 🐳 Docker

```dockerfile
# Multi-stage build
FROM node:18-alpine AS builder
# ... build stage

FROM nginx:alpine
# ... production stage
```

### 🌐 URLs de Acesso

- **Desenvolvimento**: http://localhost:5173
- **Produção (Docker)**: http://localhost:3000
- **Health Check**: http://localhost:3000/health

## 🎯 Próximos Passos

1. **🔔 WebSocket**: Notificações em tempo real
2. **🔐 Autenticação**: Login e proteção de rotas
3. **📊 Analytics**: Métricas de uso
4. **🧪 Testes**: Unit e E2E tests
5. **📱 PWA**: Progressive Web App
6. **🎨 Themes**: Dark/Light mode
7. **🌍 i18n**: Internacionalização

## 🤝 Integração com Microserviços

Este frontend é parte do sistema completo de microserviços:

- **Passo 01-07**: Microserviços Go + gRPC
- **Passo 08**: BFF GraphQL (Gateway)
- **Passo 09**: Frontend React (Este projeto)

A aplicação demonstra o fluxo completo:
**Pedido Criado** → **Pagamento Processado** → **Notificação Enviada** → **Frontend Atualizado**

## React Compiler

The React Compiler is not enabled on this template because of its impact on dev & build performances. To add it, see [this documentation](https://react.dev/learn/react-compiler/installation).

## Expanding the ESLint configuration

If you are developing a production application, we recommend using TypeScript with type-aware lint rules enabled. Check out the [TS template](https://github.com/vitejs/vite/tree/main/packages/create-vite/template-react-ts) for information on how to integrate TypeScript and [`typescript-eslint`](https://typescript-eslint.io) in your project.
