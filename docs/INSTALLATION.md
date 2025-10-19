# 📦 Guia de Instalação e Configuração

## 📋 Índice

1. [Pré-requisitos](#-pré-requisitos)
2. [Instalação Rápida](#-instalação-rápida)
3. [Instalação Detalhada](#-instalação-detalhada)
4. [Configuração](#-configuração)
5. [Dados Iniciais](#-dados-iniciais)
6. [Verificação](#-verificação)
7. [Troubleshooting](#-troubleshooting)

---

## 🔧 Pré-requisitos

### Obrigatórios

| Software | Versão Mínima | Download |
|----------|---------------|----------|
| **Docker** | 20.10+ | [docker.com](https://www.docker.com/get-started) |
| **Docker Compose** | 2.0+ | Incluído no Docker Desktop |
| **Git** | 2.30+ | [git-scm.com](https://git-scm.com/downloads) |

### Opcionais (para desenvolvimento)

| Software | Versão | Uso |
|----------|--------|-----|
| **Go** | 1.21+ | Desenvolvimento dos microserviços |
| **Node.js** | 18+ | Desenvolvimento do frontend |
| **Make** | 4.0+ | Comandos simplificados via Makefile |
| **kubectl** | 1.28+ | Deploy em Kubernetes |

### Requisitos de Sistema

- **CPU**: 4 cores (recomendado)
- **RAM**: 8GB mínimo, 16GB recomendado
- **Disco**: 10GB de espaço livre
- **SO**: Windows 10+, macOS 10.15+, ou Linux

---

## ⚡ Instalação Rápida

Para usuários experientes que querem subir tudo rapidamente:

```bash
# 1. Clone o repositório
git clone https://github.com/TaviloBreno/go-microservices-architecture.git
cd go-microservices-architecture

# 2. Configure variáveis de ambiente (opcional)
cp .env.example .env

# 3. Suba todos os serviços
docker-compose up -d

# 4. Aguarde inicialização (30-60 segundos)
sleep 60

# 5. Verifique saúde dos serviços
docker-compose ps
```

✅ **Pronto!** Acesse: http://localhost:3001

---

## 📚 Instalação Detalhada

### Passo 1: Instalar Docker

#### Windows

1. Baixe o Docker Desktop: https://www.docker.com/products/docker-desktop
2. Execute o instalador
3. Reinicie o computador
4. Abra o Docker Desktop
5. Verifique a instalação:

```powershell
docker --version
docker-compose --version
```

#### macOS

```bash
# Via Homebrew
brew install --cask docker

# Ou baixe o instalador
# https://www.docker.com/products/docker-desktop
```

#### Linux (Ubuntu/Debian)

```bash
# Adicione o repositório Docker
sudo apt-get update
sudo apt-get install ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

# Instale o Docker
echo \
  "deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  "$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Adicione seu usuário ao grupo docker
sudo usermod -aG docker $USER
newgrp docker

# Verifique
docker --version
docker compose version
```

### Passo 2: Clonar o Repositório

```bash
# Via HTTPS
git clone https://github.com/TaviloBreno/go-microservices-architecture.git

# Ou via SSH
git clone git@github.com:TaviloBreno/go-microservices-architecture.git

# Entre no diretório
cd go-microservices-architecture
```

### Passo 3: Configurar Variáveis de Ambiente (Opcional)

O projeto funciona com valores padrão, mas você pode personalizar:

```bash
# Copie o arquivo de exemplo
cp .env.example .env

# Edite conforme necessário
nano .env  # ou vim, code, notepad++
```

#### Variáveis Principais

```env
# Database
MYSQL_ROOT_PASSWORD=root123
MYSQL_DATABASE=microservices
MYSQL_USER=microservices
MYSQL_PASSWORD=micro123

# RabbitMQ
RABBITMQ_DEFAULT_USER=guest
RABBITMQ_DEFAULT_PASS=guest

# Grafana
GRAFANA_ADMIN_USER=admin
GRAFANA_ADMIN_PASSWORD=admin123

# Services
ORDER_SERVICE_PORT=50051
PAYMENT_SERVICE_PORT=50052
USER_SERVICE_PORT=50053
NOTIFICATION_SERVICE_PORT=50054
CATALOG_SERVICE_PORT=50055
BFF_PORT=8080

# Monitoring
PROMETHEUS_PORT=9090
GRAFANA_PORT=3000
JAEGER_PORT=16686
```

### Passo 4: Build das Imagens (Primeira vez)

```bash
# Build de todas as imagens
docker-compose build

# Ou build individual
docker-compose build order-service
docker-compose build payment-service
# ... etc
```

**Tempo estimado:** 5-10 minutos na primeira vez

### Passo 5: Iniciar os Serviços

```bash
# Iniciar em background
docker-compose up -d

# Ou com logs visíveis (útil para debug)
docker-compose up

# Ver logs de um serviço específico
docker-compose logs -f order-service
```

### Passo 6: Aguardar Inicialização

Os serviços levam alguns segundos para iniciar completamente:

```bash
# Verificar status
docker-compose ps

# Aguardar MySQL inicializar (importante!)
echo "Aguardando MySQL..."
sleep 30

# Verificar logs do MySQL
docker-compose logs mysql
```

### Passo 7: Verificar Saúde dos Serviços

```bash
# Usando script incluído
bash scripts/health-check.sh

# Ou via Makefile
make health-check

# Ou manualmente
curl http://localhost:8080/health       # BFF
curl http://localhost:9090/-/healthy    # Prometheus
curl http://localhost:3000/api/health   # Grafana
```

---

## ⚙️ Configuração

### Portas Utilizadas

Certifique-se de que estas portas estão livres:

| Serviço | Porta(s) | Descrição |
|---------|----------|-----------|
| MySQL | 3306 | Banco de dados |
| RabbitMQ | 5672, 15672 | Message broker + Management |
| Order Service | 50051 | gRPC |
| Payment Service | 50052 | gRPC |
| User Service | 50053 | gRPC |
| Notification Service | 50054 | gRPC |
| Catalog Service | 50055 | gRPC |
| BFF GraphQL | 8080 | API GraphQL |
| Frontend | 3001 | Interface React |
| Prometheus | 9090 | Métricas |
| Grafana | 3000 | Dashboards |
| Jaeger | 16686, 14268 | Tracing |

### Verificar Portas em Uso

#### Windows (PowerShell)
```powershell
# Verificar porta específica
netstat -ano | findstr :3306

# Liberar porta (matar processo)
taskkill /PID <PID> /F
```

#### macOS/Linux
```bash
# Verificar porta específica
lsof -i :3306

# Liberar porta
kill -9 <PID>
```

### Configurar Hosts (Opcional)

Para usar domínios locais em vez de localhost:

#### Windows
1. Abra como Administrador: `C:\Windows\System32\drivers\etc\hosts`
2. Adicione:
```
127.0.0.1 microservices.local
127.0.0.1 grafana.local
127.0.0.1 prometheus.local
127.0.0.1 jaeger.local
```

#### macOS/Linux
```bash
sudo nano /etc/hosts
```
Adicione:
```
127.0.0.1 microservices.local
127.0.0.1 grafana.local
127.0.0.1 prometheus.local
127.0.0.1 jaeger.local
```

---

## 🗄️ Dados Iniciais

O banco de dados é **automaticamente populado** na primeira inicialização com:

### Dados Criados

| Tabela | Quantidade | Descrição |
|--------|------------|-----------|
| **users** | 10 | Usuários de exemplo |
| **categories** | 8 | Categorias de produtos |
| **products** | 28 | Produtos variados |
| **orders** | 10 | Pedidos de exemplo |
| **order_items** | 15 | Itens dos pedidos |
| **payments** | 10 | Pagamentos processados |
| **notifications** | 10 | Notificações enviadas |

### Usuários de Teste

| Email | Senha | Observação |
|-------|-------|------------|
| joao.silva@email.com | password123 | Usuário com pedidos |
| maria.santos@email.com | password123 | Usuário ativo |
| pedro.oliveira@email.com | password123 | Usuário novo |

**Nota:** As senhas são criptografadas com bcrypt no banco.

### Produtos de Exemplo

- **Eletrônicos**: Smartphones, Notebooks, Fones, Smart Watches
- **Livros**: Clean Code, DDD, Design Patterns, Microservices
- **Roupas**: Camisetas, Calças, Jaquetas, Tênis
- **Casa**: Luminárias, Quadros, Jogos de Cama
- **Esportes**: Bolas, Halteres, Tapetes de Yoga
- **Alimentos**: Café, Chocolate, Azeite
- **Beleza**: Perfumes, Kits de Skin Care
- **Brinquedos**: Lego, Bonecas, Carrinhos

### Visualizar Dados

#### Via MySQL Client

```bash
# Conectar ao MySQL
docker exec -it go-microservices-architecture-mysql-1 mysql -u root -p

# Senha: root123

# Listar databases
SHOW DATABASES;

# Usar database
USE catalog_service;

# Ver produtos
SELECT id, name, price, stock_quantity FROM products LIMIT 10;

# Ver pedidos
USE order_service;
SELECT id, user_id, status, total_amount FROM orders;
```

#### Via phpMyAdmin (Opcional)

Adicione ao `docker-compose.yml`:

```yaml
phpmyadmin:
  image: phpmyadmin/phpmyadmin
  environment:
    PMA_HOST: mysql
    PMA_PORT: 3306
    MYSQL_ROOT_PASSWORD: root123
  ports:
    - "8081:80"
  depends_on:
    - mysql
```

Acesse: http://localhost:8081

---

## ✅ Verificação

### 1. Verificar Containers

```bash
# Listar containers rodando
docker-compose ps

# Deve mostrar TODOS os serviços como "Up"
```

**Saída esperada:**
```
NAME                           STATUS    PORTS
mysql                          Up        0.0.0.0:3306->3306/tcp
rabbitmq                       Up        0.0.0.0:5672->5672/tcp, 0.0.0.0:15672->15672/tcp
order-service                  Up        0.0.0.0:50051->50051/tcp
payment-service                Up        0.0.0.0:50052->50052/tcp
user-service                   Up        0.0.0.0:50053->50053/tcp
notification-service           Up        0.0.0.0:50054->50054/tcp
catalog-service                Up        0.0.0.0:50055->50055/tcp
bff                            Up        0.0.0.0:8080->8080/tcp
frontend                       Up        0.0.0.0:3001->3001/tcp
prometheus                     Up        0.0.0.0:9090->9090/tcp
grafana                        Up        0.0.0.0:3000->3000/tcp
jaeger                         Up        0.0.0.0:16686->16686/tcp
```

### 2. Testar Endpoints

```bash
# Frontend
curl http://localhost:3001

# BFF GraphQL
curl http://localhost:8080/health

# Prometheus
curl http://localhost:9090/-/healthy

# Grafana
curl http://localhost:3000/api/health

# Jaeger
curl http://localhost:16686

# RabbitMQ Management
curl http://localhost:15672
```

### 3. Testar GraphQL

```bash
# Query de teste
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ __schema { queryType { name } } }"
  }'
```

### 4. Verificar Logs

```bash
# Logs de todos os serviços
docker-compose logs

# Logs de um serviço específico
docker-compose logs -f order-service

# Últimas 100 linhas
docker-compose logs --tail=100

# Logs com timestamps
docker-compose logs -t
```

### 5. Acessar Interfaces

Abra no navegador:

- ✅ **Frontend**: http://localhost:3001
- ✅ **GraphQL Playground**: http://localhost:8080/graphql
- ✅ **Grafana**: http://localhost:3000 (admin / admin123)
- ✅ **Prometheus**: http://localhost:9090
- ✅ **Jaeger**: http://localhost:16686
- ✅ **RabbitMQ**: http://localhost:15672 (guest / guest)

---

## 🔧 Troubleshooting

### Problema: Container não inicia

**Sintomas:** Container em estado "Exit" ou "Restarting"

**Soluções:**

```bash
# Ver logs do container
docker-compose logs <service-name>

# Recriar container
docker-compose up -d --force-recreate <service-name>

# Rebuild da imagem
docker-compose build --no-cache <service-name>
docker-compose up -d <service-name>
```

### Problema: Porta já em uso

**Sintomas:** Erro "bind: address already in use"

**Soluções:**

```bash
# Windows
netstat -ano | findstr :<PORT>
taskkill /PID <PID> /F

# macOS/Linux
lsof -i :<PORT>
kill -9 <PID>

# Ou altere a porta no docker-compose.yml
```

### Problema: MySQL não conecta

**Sintomas:** Erro "Connection refused" ou "Access denied"

**Soluções:**

```bash
# Aguardar MySQL inicializar completamente
sleep 30

# Verificar logs do MySQL
docker-compose logs mysql

# Resetar volumes (CUIDADO: apaga dados!)
docker-compose down -v
docker-compose up -d
```

### Problema: Erro de memória

**Sintomas:** Containers sendo mortos (OOMKilled)

**Soluções:**

1. Aumentar memória do Docker Desktop:
   - Windows/Mac: Docker Desktop → Settings → Resources → Memory
   - Recomendado: 8GB ou mais

2. Reduzir serviços ativos:
```bash
# Subir apenas serviços essenciais
docker-compose up -d mysql rabbitmq order-service bff frontend
```

### Problema: Build falha

**Sintomas:** Erro durante `docker-compose build`

**Soluções:**

```bash
# Limpar cache do Docker
docker system prune -a

# Build sem cache
docker-compose build --no-cache

# Verificar espaço em disco
df -h  # Linux/Mac
```

### Problema: GraphQL não responde

**Sintomas:** Erro 502 ou timeout

**Soluções:**

```bash
# Verificar se microserviços estão rodando
docker-compose ps

# Verificar logs do BFF
docker-compose logs -f bff

# Reiniciar BFF
docker-compose restart bff
```

### Problema: Grafana sem dados

**Sintomas:** Dashboards vazios

**Soluções:**

1. Verificar datasource do Prometheus:
   - Grafana → Configuration → Data Sources
   - URL deve ser: http://prometheus:9090

2. Verificar se Prometheus está coletando métricas:
   - Acesse: http://localhost:9090/targets
   - Todos targets devem estar "UP"

3. Gerar tráfego para criar métricas:
```bash
# Fazer algumas requisições
for i in {1..10}; do
  curl http://localhost:8080/health
  sleep 1
done
```

### Problema: Frontend não carrega

**Sintomas:** Página em branco ou erro de build

**Soluções:**

```bash
# Ver logs do frontend
docker-compose logs -f frontend

# Rebuild do frontend
docker-compose build --no-cache frontend
docker-compose up -d frontend

# Verificar se API está acessível
curl http://localhost:8080/health
```

---

## 🧹 Comandos de Limpeza

### Parar todos os serviços

```bash
docker-compose down
```

### Parar e remover volumes (apaga dados!)

```bash
docker-compose down -v
```

### Limpar cache do Docker

```bash
# Remover images não utilizadas
docker image prune -a

# Remover volumes não utilizados
docker volume prune

# Limpeza completa
docker system prune -a --volumes
```

### Recomeçar do zero

```bash
# 1. Para e remove tudo
docker-compose down -v

# 2. Remove imagens do projeto
docker images | grep go-ms | awk '{print $3}' | xargs docker rmi -f

# 3. Rebuild completo
docker-compose build --no-cache

# 4. Sobe novamente
docker-compose up -d
```

---

## 📖 Próximos Passos

Após a instalação bem-sucedida:

1. ✅ **Explore o Frontend**: http://localhost:3001
2. ✅ **Teste a API GraphQL**: http://localhost:8080/graphql
3. ✅ **Configure Dashboards**: http://localhost:3000
4. ✅ **Veja Traces**: http://localhost:16686
5. ✅ **Leia a Documentação**: [docs/](../docs/)

### Documentação Adicional

- [Arquitetura do Sistema](ARCHITECTURE.md)
- [Guia de Desenvolvimento](QUICKSTART.md)
- [Monitoramento](PASSO-10-MONITORING.md)
- [CI/CD](PASSO-11-CICD.md)
- [Deploy em Produção](PASSO-12-DEPLOYMENT.md)

---

## 🆘 Precisa de Ajuda?

- 📖 **Documentação**: Ver pasta `docs/`
- 🐛 **Reportar Bug**: [GitHub Issues](https://github.com/TaviloBreno/go-microservices-architecture/issues)
- 💬 **Discussões**: [GitHub Discussions](https://github.com/TaviloBreno/go-microservices-architecture/discussions)

---

**Última Atualização:** Outubro 2025  
**Versão:** 1.0.0
