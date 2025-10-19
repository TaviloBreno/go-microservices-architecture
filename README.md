# 🧠 Go Microservices Architecture

[![CI Status](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/ci.yml/badge.svg)](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/ci.yml)
[![CD Status](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/cd.yml/badge.svg)](https://github.com/TaviloBreno/go-microservices-architecture/actions/workflows/cd.yml)
![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?logo=docker)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

Arquitetura moderna e escalável utilizando:
- **Go (Golang)** — gRPC, GraphQL, Clean Architecture
- **RabbitMQ** — mensageria assíncrona
- **MySQL** — persistência por microsserviço
- **React** — frontend SPA
- **Docker Compose / K8s** — infraestrutura completa
- **OpenTelemetry + Prometheus + Jaeger** — observabilidade
- **GitHub Actions** — CI/CD automatizado

---

## 🚀 Estrutura do projeto


## ⚙️ Tecnologias

| Categoria | Tecnologia |
|------------|-------------|
| Backend | Go (1.22+), gRPC, GraphQL, Clean Architecture |
| Mensageria | RabbitMQ |
| Banco de Dados | MySQL |
| Frontend | React (Vite) |
| Infraestrutura | Docker Compose |
| Observabilidade | OpenTelemetry, Prometheus, Jaeger |

---

## 🧩 Próximos Passos

1️⃣ Criar estrutura de módulos Go  
2️⃣ Configurar Docker Compose (MySQL, RabbitMQ, Jaeger)  
3️⃣ Implementar microsserviços base com gRPC e GraphQL  

---

## 🧑‍💻 Autor

Desenvolvido por [Seu Nome]  
📧 [seuemail@example.com]  
🔗 GitHub: [https://github.com/seu-usuario](https://github.com/seu-usuario)
✅ 7️⃣ Configurar branches
bash
Copiar código
git checkout -b develop
git push origin develop
Branches padrão:

main → produção

develop → desenvolvimento

feature/* → novas features

fix/* → correções

✅ 8️⃣ Primeiro commit
bash
Copiar código
git add .
git commit -m "chore: estrutura inicial do repositório"
git push origin develop
✅ 9️⃣ (Opcional) Adicionar CI básico
Crie .github/workflows/go.yml:

yaml
Copiar código
name: Go CI

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: 1.22
      - run: go mod tidy
      - run: go build ./...
✅ 10️⃣ Resultado final esperado
Após esse passo, o repositório deve conter:

go
Copiar código
go-microservices-architecture/
│
├── .github/workflows/go.yml
├── .gitignore
├── LICENSE
├── README.md
├── docker-compose.yml
├── infra/
├── bff/
├── frontend/
└── services/
    ├── user/
    ├── catalog/
    ├── order/
    ├── payment/
    └── notification/