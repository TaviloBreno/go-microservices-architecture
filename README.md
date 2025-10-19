# 🧠 Go Microservices Architecture

[![CI Status](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/ci.yml/badge.svg)](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/ci.yml)
[![CD Status](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/cd.yml/badge.svg)](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/cd.yml)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

> Arquitetura completa de microserviços em Go com gRPC, GraphQL, monitoramento distribuído e CI/CD automatizado.

## 📋 Índice

- [Sobre o Projeto](#-sobre-o-projeto)
- [Arquitetura](#-arquitetura)
- [Stack Tecnológica](#-stack-tecnológica)
- [Início Rápido](#-início-rápido)
- [Documentação](#-documentação)
- [Testes e CI/CD](#-testes-e-cicd)
- [Contribuindo](#-contribuindo)

---

## 🎯 Sobre o Projeto

Este projeto demonstra uma arquitetura moderna e escalável de microserviços utilizando as melhores práticas de desenvolvimento em Go, incluindo:

✅ **6 Microserviços independentes** (Order, Payment, User, Notification, Catalog, BFF)  
✅ **Comunicação gRPC** de alta performance  
✅ **API GraphQL** via Backend for Frontend (BFF)  
✅ **Message Queue** com RabbitMQ para processamento assíncrono  
✅ **Observabilidade completa** com Prometheus, Grafana e Jaeger  
✅ **Testes automatizados** com 70%+ de cobertura  
✅ **CI/CD** com GitHub Actions  
✅ **Dashboard React** com suporte a dark mode  
✅ **Containerização completa** com Docker Compose  

---

## 🏗️ Arquitetura

```
Frontend (React) → BFF (GraphQL) → Microservices (gRPC)
                                    ↓
                        MySQL + RabbitMQ + Observability
```

**Microserviços:**
- 🛒 **Order Service** (50051) - Gestão de pedidos
- 💳 **Payment Service** (50052) - Processamento de pagamentos
- 👤 **User Service** (50053) - Gestão de usuários
- 📧 **Notification Service** (50054) - Envio de notificações
- 📦 **Catalog Service** (50055) - Catálogo de produtos
- 🔷 **BFF GraphQL** (8080) - API unificada

**Infraestrutura:**
- 🐬 **MySQL** (3306) - Banco de dados relacional
- 🐰 **RabbitMQ** (5672, 15672) - Message broker
- 📊 **Prometheus** (9090) - Coleta de métricas
- 📈 **Grafana** (3000) - Visualização e dashboards
- � **Jaeger** (16686) - Distributed tracing

📖 **[Documentação Completa da Arquitetura](docs/ARCHITECTURE.md)**

---

## ⚙️ Stack Tecnológica

| Categoria | Tecnologia | Versão |
|-----------|-----------|--------|
| **Backend** | Go | 1.21+ |
| **RPC** | gRPC | - |
| **API** | GraphQL | - |
| **Database** | MySQL | 8.0+ |
| **Message Queue** | RabbitMQ | 3.12+ |
| **Frontend** | React | 18+ |
| **Metrics** | Prometheus | 2.48+ |
| **Visualization** | Grafana | 10.2+ |
| **Tracing** | Jaeger | 1.55+ |
| **Container** | Docker | 20.10+ |
| **Orchestration** | Docker Compose | 2.0+ |
| **CI/CD** | GitHub Actions | - |
| **Testing** | testify | 1.9.0 |

---

## 🚀 Início Rápido

### Pré-requisitos

- Docker 20.10+ e Docker Compose 2.0+
- Go 1.21+ (para desenvolvimento local)
- Node.js 18+ (para frontend)

### Subir toda a infraestrutura

```bash
# Clonar o repositório
git clone <repository-url>
cd go-microservices-architecture

# Subir todos os serviços
docker-compose up -d

# Ou usando Makefile (recomendado)
make docker-up

# Verificar saúde dos serviços
make health-check
```

### Acessar os serviços

| Serviço | URL | Credenciais |
|---------|-----|-------------|
| 🎨 Frontend | http://localhost:3001 | - |
| � GraphQL | http://localhost:8080/graphql | - |
| 📊 Grafana | http://localhost:3000 | admin / admin123 |
| 📈 Prometheus | http://localhost:9090 | - |
| 🔍 Jaeger | http://localhost:16686 | - |
| 🐰 RabbitMQ | http://localhost:15672 | guest / guest |

📖 **[Guia de Início Rápido Completo](docs/QUICKSTART.md)**

---

## 📚 Documentação

| Documento | Descrição |
|-----------|-----------|
| [🏛️ Arquitetura](docs/ARCHITECTURE.md) | Arquitetura completa do sistema |
| [🚀 Início Rápido](docs/QUICKSTART.md) | Guia para começar em minutos |
| [📊 Passo 10: Monitoring](docs/PASSO-10-MONITORING.md) | Prometheus, Grafana e Jaeger |
| [🔄 Passo 11: CI/CD](docs/PASSO-11-CICD.md) | GitHub Actions e testes |

---

## 🧪 Testes e CI/CD

### Executar todos os testes

```bash
# Via script completo
bash scripts/run-all-tests.sh

# Ou via Makefile
make test

# Simular CI localmente
make ci-local
```

### Testes por serviço

```bash
make test-order      # Order Service
make test-payment    # Payment Service
make test-user       # User Service
make test-frontend   # React Frontend
```

### Cobertura de código

```bash
make coverage
```

### CI/CD Pipeline

- ✅ **CI**: Testes automáticos em cada push/PR
- ✅ **CD**: Deploy automático em tags (v*.*.*)
- ✅ **Linting**: golangci-lint com 15+ linters
- ✅ **Coverage**: Relatório de cobertura no CI
- ✅ **Docker**: Build multi-platform (amd64, arm64)

📖 **[Documentação Completa de CI/CD](docs/PASSO-11-CICD.md)**

---

## 🛠️ Comandos Úteis (Makefile)

```bash
make help             # Mostra todos os comandos disponíveis
make docker-up        # Sobe todos os containers
make docker-down      # Para todos os containers
make health-check     # Verifica saúde dos serviços
make test             # Roda todos os testes
make lint             # Executa linter
make coverage         # Gera relatório de cobertura
make prometheus       # Abre Prometheus no browser
make grafana          # Abre Grafana no browser
make jaeger           # Abre Jaeger no browser
make clean            # Limpa arquivos temporários
```

---

## 📁 Estrutura do Projeto

```
go-microservices-architecture/
├── .github/
│   └── workflows/          # CI/CD pipelines
│       ├── ci.yml          # Continuous Integration
│       └── cd.yml          # Continuous Deployment
├── bff/                    # Backend for Frontend (GraphQL)
│   ├── cmd/
│   │   └── main.go
│   ├── Dockerfile
│   └── go.mod
├── services/               # Microserviços
│   ├── order/              # Order Service
│   │   ├── cmd/
│   │   ├── internal/
│   │   │   ├── domain/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   ├── handler/
│   │   │   ├── telemetry/
│   │   │   └── metrics/
│   │   ├── Dockerfile
│   │   └── go.mod
│   ├── payment/            # Payment Service
│   ├── user/               # User Service
│   ├── notification/       # Notification Service
│   └── catalog/            # Catalog Service
├── frontend/               # React Dashboard
│   ├── src/
│   ├── package.json
│   └── vite.config.js
├── infra/                  # Infraestrutura
│   ├── mysql/
│   │   └── init/
│   └── prometheus/
│       └── prometheus.yml
├── docs/                   # Documentação
│   ├── ARCHITECTURE.md
│   ├── QUICKSTART.md
│   ├── PASSO-10-MONITORING.md
│   └── PASSO-11-CICD.md
├── scripts/                # Scripts utilitários
│   ├── run-all-tests.sh
│   └── health-check.sh
├── .golangci.yml           # Configuração do linter
├── docker-compose.yml      # Orquestração dos serviços
├── Makefile                # Comandos úteis
└── README.md
```

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Para contribuir:

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

### Diretrizes de Contribuição

- Escreva testes para novas funcionalidades
- Mantenha a cobertura de código acima de 70%
- Siga os padrões de código (use `make lint`)
- Atualize a documentação conforme necessário
- Use Conventional Commits para mensagens de commit

---

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

## 🎓 Aprendizados

Este projeto demonstra:

- ✅ Arquitetura de microserviços escalável
- ✅ Comunicação gRPC de alta performance
- ✅ API GraphQL moderna
- ✅ Observabilidade distribuída
- ✅ Testes automatizados
- ✅ CI/CD com GitHub Actions
- ✅ Containerização com Docker
- ✅ Clean Architecture em Go
- ✅ Message Queue com RabbitMQ
- ✅ Monitoramento em tempo real

---

## 🆘 Suporte

- 📖 **Documentação**: Ver pasta `docs/`
- 🐛 **Issues**: Reporte bugs via GitHub Issues
- 💬 **Discussões**: Use GitHub Discussions
- 📧 **Email**: [seu-email@example.com]

---

## 🙏 Agradecimentos

- [Go Team](https://golang.org/) pela linguagem incrível
- [gRPC Team](https://grpc.io/) pelo framework de RPC
- [GraphQL Foundation](https://graphql.org/) pela especificação
- [Prometheus](https://prometheus.io/) e [Grafana](https://grafana.com/) pela stack de observabilidade
- [Jaeger](https://www.jaegertracing.io/) pelo distributed tracing
- Comunidade Go pela inspiração e recursos

---

**Desenvolvido com ❤️ usando Go**

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-00ADD8?logo=go)](https://golang.org/)
[![Powered by gRPC](https://img.shields.io/badge/Powered%20by-gRPC-244c5a?logo=grpc)](https://grpc.io/)
[![Built with Docker](https://img.shields.io/badge/Built%20with-Docker-2496ED?logo=docker)](https://www.docker.com/)

└── services/
    ├── user/
    ├── catalog/
    ├── order/
    ├── payment/
    └── notification/