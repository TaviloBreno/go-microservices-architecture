# 🚀 Passo 11 - CI/CD com GitHub Actions

Este passo implementa um sistema completo de integração e entrega contínua (CI/CD) usando GitHub Actions, com testes automatizados para todos os microsserviços.

## 🎯 Objetivos Alcançados

- ✅ Testes unitários automatizados em todos os microsserviços
- ✅ Pipeline de CI que roda em cada PR e push
- ✅ Pipeline de CD que faz deploy automático em releases
- ✅ Linting e análise estática de código
- ✅ Cobertura de código e relatórios
- ✅ Build e push automático de imagens Docker
- ✅ Validação de Docker Compose
- ✅ Scripts para testes locais

## 📊 Arquitetura de CI/CD

```
┌─────────────┐
│   Push/PR   │
└──────┬──────┘
       │
       ▼
┌─────────────────────────────────────┐
│      GitHub Actions (CI)            │
├─────────────────────────────────────┤
│  1. Test Go Services (6 jobs)       │
│  2. Test BFF GraphQL                │
│  3. Lint All Services               │
│  4. Test Frontend React             │
│  5. Docker Build Validation         │
│  6. Docker Compose Validation       │
└──────┬──────────────────────────────┘
       │
       ▼
┌─────────────────┐
│  All Tests Pass │
└──────┬──────────┘
       │
       ▼
┌─────────────────────────────────────┐
│   Tag Release (v*.*.*)              │
└──────┬──────────────────────────────┘
       │
       ▼
┌─────────────────────────────────────┐
│      GitHub Actions (CD)            │
├─────────────────────────────────────┤
│  1. Build Docker Images (7 images)  │
│  2. Push to Docker Hub              │
│  3. Create GitHub Release           │
│  4. Send Notifications              │
└─────────────────────────────────────┘
```

## 🧪 Estrutura de Testes

### Testes Criados

#### Order Service
📁 `services/order/internal/service/order_service_test.go`
- ✅ TestCreateOrder_Success
- ✅ TestCreateOrder_RepositoryFailure
- ✅ TestFindOrderByID_Success
- ✅ TestFindOrderByID_NotFound
- ✅ TestUpdateOrder_Success
- ✅ TestPublishOrderCreated_Success
- ✅ BenchmarkCreateOrder
- ✅ TestCreateOrderWithContext_Timeout
- ✅ TestOrderOperations_Parallel

#### Payment Service
📁 `services/payment/internal/service/payment_service_test.go`
- ✅ TestProcessPayment_Success
- ✅ TestProcessPayment_InvalidAmount
- ✅ TestFindPaymentByOrderID_Success
- ✅ TestUpdatePaymentStatus_Success
- ✅ TestValidatePaymentMethod
- ✅ TestCalculatePaymentFee
- ✅ BenchmarkProcessPayment

#### User Service
📁 `services/user/internal/service/user_service_test.go`
- ✅ TestCreateUser_Success
- ✅ TestCreateUser_DuplicateEmail
- ✅ TestValidateEmail
- ✅ TestFindUserByID_Success
- ✅ TestFindUserByEmail_Success
- ✅ TestUpdateUser_Success
- ✅ TestDeleteUser_Success
- ✅ BenchmarkCreateUser

## 📋 Workflows GitHub Actions

### 1. CI Pipeline (`.github/workflows/ci.yml`)

**Triggers:**
- Push em `main` ou `develop`
- Pull Requests para `main` ou `develop`

**Jobs:**

#### test-go-services
- Roda em paralelo para: order, payment, user, notification, catalog
- Executa: `go test ./... -v -race -coverprofile=coverage.out`
- Gera relatório de cobertura
- Upload para Codecov

#### test-bff
- Testa o BFF GraphQL separadamente
- Cobertura de código completa

#### lint-go
- Executa golangci-lint em todos os serviços
- Verifica código seguindo best practices

#### test-frontend
- Instala dependências Node.js
- Roda ESLint
- Faz build do React

#### docker-build-test
- Valida que todas as imagens Docker podem ser buildadas
- Usa cache para otimizar tempo de build

#### docker-compose-test
- Valida sintaxe do docker-compose.yml

### 2. CD Pipeline (`.github/workflows/cd.yml`)

**Triggers:**
- Tags no formato `v*.*.*` (ex: v1.0.0, v2.1.5)
- Dispatch manual via interface do GitHub

**Jobs:**

#### build-and-push
- Builda imagens Docker para todos os serviços
- Suporta múltiplas plataformas (amd64, arm64)
- Faz push para Docker Hub
- Tags: latest, semver, sha

#### create-release
- Gera changelog automaticamente
- Cria release no GitHub
- Documenta como usar as imagens Docker

#### notify
- Envia sumário do deployment
- Status de cada serviço

## 🔧 Configuração Necessária

### Secrets do GitHub

Configure em: **Settings → Secrets and variables → Actions**

```bash
DOCKERHUB_USERNAME=seu-usuario-dockerhub
DOCKERHUB_TOKEN=seu-token-dockerhub
GITHUB_TOKEN=<gerado automaticamente>
```

### Como obter Docker Hub Token:
1. Acesse https://hub.docker.com/settings/security
2. Clique em "New Access Token"
3. Dê um nome descritivo (ex: "github-actions")
4. Copie o token gerado

## 🚀 Como Usar

### 1. Desenvolvimento Local

Antes de fazer push, rode os testes localmente:

```bash
# Linux/Mac
chmod +x scripts/run-all-tests.sh
./scripts/run-all-tests.sh

# Windows (Git Bash)
bash scripts/run-all-tests.sh
```

### 2. Criar Pull Request

```bash
git checkout -b feature/nova-funcionalidade
# Faça suas alterações
git add .
git commit -m "feat: adiciona nova funcionalidade"
git push origin feature/nova-funcionalidade
```

O CI será executado automaticamente! Verifique em:
- **Actions** tab no GitHub
- Você verá todos os checks rodando

### 3. Fazer Release

```bash
# Atualizar versão e criar tag
git tag v1.0.0
git push origin v1.0.0
```

O CD será executado automaticamente:
1. ✅ Builda todas as imagens Docker
2. ✅ Faz push para Docker Hub
3. ✅ Cria release no GitHub com changelog
4. ✅ Notifica sobre o deployment

### 4. Deploy Manual

Você também pode fazer deploy manual:

1. Vá em **Actions** → **CD Pipeline**
2. Clique em "Run workflow"
3. Digite a tag version (ex: v1.0.0)
4. Clique em "Run workflow"

## 📊 Badges de Status

Adicione ao README.md:

```markdown
![CI Status](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/ci.yml/badge.svg)
![CD Status](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/cd.yml/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
```

## 🐳 Imagens Docker Geradas

Após cada release, as seguintes imagens estarão disponíveis:

```bash
# Order Service
docker pull seu-usuario/go-ms-order:latest
docker pull seu-usuario/go-ms-order:v1.0.0

# Payment Service  
docker pull seu-usuario/go-ms-payment:latest
docker pull seu-usuario/go-ms-payment:v1.0.0

# User Service
docker pull seu-usuario/go-ms-user:latest
docker pull seu-usuario/go-ms-user:v1.0.0

# Notification Service
docker pull seu-usuario/go-ms-notification:latest
docker pull seu-usuario/go-ms-notification:v1.0.0

# Catalog Service
docker pull seu-usuario/go-ms-catalog:latest
docker pull seu-usuario/go-ms-catalog:v1.0.0

# BFF GraphQL
docker pull seu-usuario/go-ms-bff-graphql:latest
docker pull seu-usuario/go-ms-bff-graphql:v1.0.0

# Frontend
docker pull seu-usuario/go-ms-frontend:latest
docker pull seu-usuario/go-ms-frontend:v1.0.0
```

## 📈 Cobertura de Código

A cobertura de código é calculada automaticamente e enviada para o Codecov:

```bash
# Ver cobertura localmente
cd services/order
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 🔍 Linting

### Executar Localmente

```bash
# Instalar golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Rodar em um serviço
cd services/order
golangci-lint run

# Rodar em todos os serviços
for service in order payment user notification catalog bff-graphql; do
  echo "Linting $service..."
  cd services/$service
  golangci-lint run
  cd ../..
done
```

### Configuração

O arquivo `.golangci.yml` na raiz do projeto configura:
- Linters habilitados
- Regras de exclusão
- Limites de complexidade
- Formatação de output

## 🧪 Executar Testes Específicos

```bash
# Testar um serviço específico
cd services/order
go test ./... -v

# Testar com cobertura
go test ./... -v -coverprofile=coverage.out

# Testar com race detector
go test ./... -race

# Rodar apenas um teste específico
go test -run TestCreateOrder_Success

# Benchmarks
go test -bench=. -benchmem
```

## 🎯 Próximos Passos

1. **Testes de Integração**: Adicionar testes E2E
2. **Testes de Carga**: Integrar k6 ou Artillery
3. **Security Scanning**: Adicionar Snyk ou Trivy
4. **Dependabot**: Automatizar updates de dependências
5. **Staging Environment**: Deploy automático para staging
6. **Rollback Automático**: Implementar rollback em caso de falha

## 🚨 Troubleshooting

### Tests falham localmente mas passam no CI
- Verificar versão do Go (deve ser 1.21+)
- Limpar cache: `go clean -cache -testcache`
- Verificar dependências: `go mod verify`

### Docker build falha
- Verificar Dockerfile
- Testar build local: `docker build -t test ./services/order`
- Verificar logs no GitHub Actions

### Push para Docker Hub falha
- Verificar se secrets estão configurados
- Verificar se token não expirou
- Verificar se usuário tem permissões

---

✅ **Sistema de CI/CD Completo Implementado!**

Agora todos os microsserviços têm testes automatizados, linting, e deployment automático! 🎉