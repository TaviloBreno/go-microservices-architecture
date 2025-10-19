# 🚀 Passo 12: Deploy Automatizado em Produção

## 📋 Índice

1. [Visão Geral](#-visão-geral)
2. [Opções de Deploy](#-opções-de-deploy)
3. [Docker Swarm](#-docker-swarm)
4. [Kubernetes](#-kubernetes)
5. [CI/CD Automatizado](#-cicd-automatizado)
6. [Secrets e Segurança](#-secrets-e-segurança)
7. [Monitoramento em Produção](#-monitoramento-em-produção)
8. [Rollback e Disaster Recovery](#-rollback-e-disaster-recovery)
9. [Troubleshooting](#-troubleshooting)

---

## 🎯 Visão Geral

O Passo 12 implementa deploy automatizado completo para produção com:

✅ **Duas opções de orquestração**: Docker Swarm e Kubernetes  
✅ **CI/CD totalmente automatizado** com GitHub Actions  
✅ **Rollback automático** em caso de falhas  
✅ **Secrets seguros** e gestão de credenciais  
✅ **High Availability** (HA) com múltiplas réplicas  
✅ **Auto-scaling** (HPA no Kubernetes)  
✅ **Health checks** automatizados  
✅ **Zero-downtime deployments**  

### Fluxo Completo

```
┌─────────────┐
│   Commit    │
│   to main   │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  CI Tests   │  ← Testes automatizados
│  & Build    │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│  Create Tag │  ← v1.0.0, v1.1.0, etc
│  (Release)  │
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ Build Images│  ← Multi-platform builds
│ Push to Hub │
└──────┬──────┘
       │
       ├──────────────┬──────────────┐
       ▼              ▼              ▼
┌─────────────┐ ┌──────────┐ ┌─────────────┐
│Docker Swarm │ │   OR    │ │ Kubernetes  │
│   Deploy    │ │         │ │   Deploy    │
└──────┬──────┘ └──────────┘ └──────┬──────┘
       │                            │
       ▼                            ▼
┌─────────────┐            ┌─────────────┐
│Health Checks│            │Health Checks│
└──────┬──────┘            └──────┬──────┘
       │                            │
       ├────────────────────────────┤
       ▼                            ▼
┌─────────────────────────────────────┐
│     Production Environment          │
│  ✅ All Services Running            │
└─────────────────────────────────────┘
```

---

## 🐳 Opções de Deploy

### Docker Swarm vs Kubernetes

| Característica | Docker Swarm | Kubernetes |
|----------------|--------------|------------|
| **Complexidade** | Simples | Avançada |
| **Setup** | Rápido (minutos) | Moderado (horas) |
| **Escalabilidade** | Boa (centenas de containers) | Excelente (milhares de containers) |
| **Auto-scaling** | Manual | Automático (HPA) |
| **Ecosistema** | Menor | Maior (Helm, Operators, etc) |
| **Curva de Aprendizado** | Baixa | Alta |
| **Melhor para** | Projetos médios | Projetos enterprise |
| **Nosso Uso** | ✅ Produção simples | ✅ Produção escalável |

**Recomendação:**
- **Docker Swarm**: Para equipes pequenas, deployments rápidos, menos infraestrutura
- **Kubernetes**: Para produção enterprise, alta escala, auto-scaling avançado

---

## 🐝 Docker Swarm

### Arquitetura

```
                 ┌──────────────────┐
                 │  Swarm Manager   │
                 │  (Leader)        │
                 └────────┬─────────┘
                          │
         ├────────────────┼────────────────┤
         ▼                ▼                ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│Worker Node 1│  │Worker Node 2│  │Worker Node 3│
│             │  │             │  │             │
│ order x3    │  │ payment x3  │  │ user x2     │
│ bff x3      │  │ catalog x2  │  │ notif x2    │
└─────────────┘  └─────────────┘  └─────────────┘
```

### 1. Inicializar Docker Swarm

```bash
# No servidor manager
docker swarm init --advertise-addr <MANAGER-IP>

# Adicionar workers (executar nos nodes workers)
docker swarm join --token <TOKEN> <MANAGER-IP>:2377

# Verificar nodes
docker node ls
```

### 2. Configurar Secrets

```bash
# Criar secrets necessários
echo "root123" | docker secret create mysql_root_password -
echo "micro123" | docker secret create mysql_password -
echo "guest123" | docker secret create rabbitmq_password -
echo "admin123" | docker secret create grafana_password -
echo "smtp_password" | docker secret create smtp_password -

# Listar secrets
docker secret ls
```

### 3. Deploy com Script

```bash
# Usar o script automatizado
cd deployment/scripts
chmod +x deploy-swarm.sh
./deploy-swarm.sh

# Ou manualmente
export DOCKERHUB_USERNAME=tavilobreno
export VERSION=v1.0.0
docker stack deploy -c deployment/docker-swarm/stack.yml --with-registry-auth go-ms
```

### 4. Verificar Deploy

```bash
# Status do stack
docker stack services go-ms

# Logs de um serviço
docker service logs go-ms_order-service

# Escalar serviço
docker service scale go-ms_order-service=5

# Health check
./deployment/scripts/health-check-swarm.sh
```

### 5. Atualizar Serviços

```bash
# Update com zero downtime
docker service update --image tavilobreno/go-ms-order:v1.1.0 go-ms_order-service

# Rollback se necessário
docker service rollback go-ms_order-service
```

### Stack Configuration

O arquivo `deployment/docker-swarm/stack.yml` inclui:

- ✅ **7 microserviços** (order, payment, user, notification, catalog, bff, frontend)
- ✅ **Infraestrutura** (MySQL, RabbitMQ)
- ✅ **Monitoring** (Prometheus, Grafana, Jaeger)
- ✅ **Networks** separadas (microservices, monitoring, frontend)
- ✅ **Secrets** para credenciais sensíveis
- ✅ **Health checks** para todos os serviços
- ✅ **Resource limits** (CPU, Memory)
- ✅ **Rolling updates** com rollback automático
- ✅ **Múltiplas réplicas** para HA

---

## ☸️ Kubernetes

### Arquitetura

```
                    ┌─────────────────┐
                    │  Control Plane  │
                    │  (Master)       │
                    └────────┬────────┘
                             │
            ├────────────────┼────────────────┤
            ▼                ▼                ▼
   ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
   │   Node 1    │  │   Node 2    │  │   Node 3    │
   │             │  │             │  │             │
   │ Pods (3-5)  │  │ Pods (3-5)  │  │ Pods (3-5)  │
   └─────────────┘  └─────────────┘  └─────────────┘
            │                │                │
            └────────────────┴────────────────┘
                             │
                    ┌────────┴────────┐
                    │  Load Balancer  │
                    │  (Ingress)      │
                    └─────────────────┘
```

### 1. Pré-requisitos

```bash
# Instalar kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# Verificar versão
kubectl version --client

# Configurar cluster (exemplo: minikube, EKS, GKE, AKS)
# Minikube (local)
minikube start --cpus=4 --memory=8192

# AWS EKS
aws eks update-kubeconfig --region us-east-1 --name my-cluster

# Google GKE
gcloud container clusters get-credentials my-cluster --zone us-central1-a
```

### 2. Deploy com Script

```bash
# Usar script automatizado
cd deployment/scripts
chmod +x deploy-k8s.sh
./deploy-k8s.sh

# O script vai:
# 1. Criar namespace
# 2. Configurar secrets
# 3. Aplicar ConfigMaps
# 4. Deploy infraestrutura
# 5. Deploy microserviços
# 6. Deploy monitoring
# 7. Configurar Ingress
```

### 3. Deploy Manual (Passo a Passo)

```bash
# 1. Namespace
kubectl apply -f deployment/kubernetes/00-namespace.yaml

# 2. Secrets
kubectl create secret generic mysql-secret -n microservices \
  --from-literal=root-password=root123 \
  --from-literal=user=microservices \
  --from-literal=password=micro123 \
  --from-literal=database=microservices

kubectl create secret generic rabbitmq-secret -n microservices \
  --from-literal=user=guest \
  --from-literal=password=guest123 \
  --from-literal=url=amqp://guest:guest123@rabbitmq:5672/

# 3. ConfigMaps
kubectl apply -f deployment/kubernetes/02-configmaps.yaml

# 4. Infraestrutura
kubectl apply -f deployment/kubernetes/30-infrastructure.yaml

# 5. Microserviços
kubectl apply -f deployment/kubernetes/10-order-service.yaml
kubectl apply -f deployment/kubernetes/11-payment-service.yaml
kubectl apply -f deployment/kubernetes/12-other-services.yaml

# 6. BFF e Frontend
kubectl apply -f deployment/kubernetes/20-bff-frontend.yaml

# 7. Monitoring
kubectl apply -f deployment/kubernetes/40-monitoring.yaml

# 8. Ingress e HPA
kubectl apply -f deployment/kubernetes/50-ingress.yaml
```

### 4. Verificar Deploy

```bash
# Status dos deployments
kubectl get deployments -n microservices

# Status dos pods
kubectl get pods -n microservices

# Logs de um pod
kubectl logs <pod-name> -n microservices

# Health check
./deployment/scripts/health-check-k8s.sh
```

### 5. Gerenciamento

```bash
# Escalar deployment
kubectl scale deployment order-service --replicas=5 -n microservices

# Update de imagem
kubectl set image deployment/order-service \
  order-service=tavilobreno/go-ms-order:v1.1.0 -n microservices

# Rollback
kubectl rollout undo deployment/order-service -n microservices

# Histórico de rollout
kubectl rollout history deployment/order-service -n microservices
```

### Kubernetes Resources

Os manifests incluem:

- ✅ **Namespace** isolado
- ✅ **Secrets** para credenciais
- ✅ **ConfigMaps** para configurações
- ✅ **Deployments** com rolling updates
- ✅ **Services** (ClusterIP, LoadBalancer)
- ✅ **StatefulSets** (MySQL, RabbitMQ)
- ✅ **Ingress** para roteamento externo
- ✅ **HPA** (Horizontal Pod Autoscaler)
- ✅ **NetworkPolicies** para segurança
- ✅ **Resource Limits** (CPU, Memory)

---

## 🔄 CI/CD Automatizado

### Workflows Implementados

#### 1. Deploy to Docker Swarm (`.github/workflows/deploy-swarm.yml`)

**Trigger:**
- Push de tag `v*.*.*` (ex: v1.0.0)
- Manual via `workflow_dispatch`

**Jobs:**
1. **pre-deploy-checks**: Valida configurações
2. **build-and-push**: Build multi-platform e push para Docker Hub
3. **deploy-swarm**: Deploy via SSH no Swarm Manager
4. **health-check**: Verificação pós-deploy
5. **rollback**: Automático se falhar
6. **notify**: Notificação do resultado

**Exemplo de uso:**
```bash
# Criar tag e fazer push (dispara deploy automaticamente)
git tag v1.0.0
git push origin v1.0.0

# Acompanhar no GitHub Actions
# https://github.com/user/repo/actions
```

#### 2. Deploy to Kubernetes (`.github/workflows/deploy-kubernetes.yml`)

**Trigger:**
- Push de tag `v*.*.*`
- Manual via `workflow_dispatch`

**Jobs:**
1. **build-and-push**: Build e push para GHCR
2. **deploy-k8s**: Deploy no cluster Kubernetes
3. **health-check**: Validação de pods
4. **rollback**: Automático em falhas
5. **notify**: Resumo do deployment

---

## 🔐 Secrets e Segurança

### Secrets Necessários

Configure no GitHub (Settings → Secrets → Actions):

#### Para Docker Swarm:

| Secret | Descrição | Exemplo |
|--------|-----------|---------|
| `DOCKERHUB_USERNAME` | Nome de usuário Docker Hub | `tavilobreno` |
| `DOCKERHUB_TOKEN` | Token de acesso | `dckr_pat_xxx...` |
| `SWARM_HOST` | IP/hostname do Swarm Manager | `192.168.1.100` |
| `SSH_USER` | Usuário SSH | `ubuntu` |
| `SSH_PRIVATE_KEY` | Chave privada SSH | `-----BEGIN RSA...` |
| `MYSQL_ROOT_PASSWORD` | Senha root MySQL | `root123` |
| `MYSQL_PASSWORD` | Senha user MySQL | `micro123` |
| `RABBITMQ_PASSWORD` | Senha RabbitMQ | `guest123` |
| `GRAFANA_PASSWORD` | Senha admin Grafana | `admin123` |
| `SMTP_PASSWORD` | Senha SMTP (opcional) | `smtp_pass` |

#### Para Kubernetes:

| Secret | Descrição |
|--------|-----------|
| `KUBECONFIG` | Config do cluster (base64) |
| `MYSQL_ROOT_PASSWORD` | Senha root MySQL |
| `MYSQL_PASSWORD` | Senha user MySQL |
| `RABBITMQ_PASSWORD` | Senha RabbitMQ |
| `GRAFANA_PASSWORD` | Senha Grafana |
| `SMTP_PASSWORD` | Senha SMTP |

### Gerar SSH Key (para Swarm)

```bash
# Gerar par de chaves
ssh-keygen -t rsa -b 4096 -C "deploy@microservices" -f ~/.ssh/deploy_key

# Copiar chave pública para servidor
ssh-copy-id -i ~/.ssh/deploy_key.pub user@swarm-host

# Copiar chave privada para GitHub Secret
cat ~/.ssh/deploy_key
# Cole todo o conteúdo em SSH_PRIVATE_KEY
```

### Gerar KUBECONFIG (para K8s)

```bash
# Obter kubeconfig atual
cat ~/.kube/config | base64 -w 0

# Cole o resultado em GitHub Secret: KUBECONFIG
```

---

## 📊 Monitoramento em Produção

### Endpoints de Monitoramento

| Serviço | URL | Credenciais |
|---------|-----|-------------|
| **Grafana** | http://localhost:3000 ou http://grafana.microservices.local | admin / (GRAFANA_PASSWORD) |
| **Prometheus** | http://localhost:9090 ou http://prometheus.microservices.local | - |
| **Jaeger** | http://localhost:16686 ou http://jaeger.microservices.local | - |

### Métricas Importantes

```promql
# Taxa de requisições
rate(grpc_server_handled_total[5m])

# Latência P95
histogram_quantile(0.95, rate(grpc_server_handling_seconds_bucket[5m]))

# Taxa de erro
rate(grpc_server_handled_total{grpc_code!="OK"}[5m])

# Uso de CPU por serviço
container_cpu_usage_seconds_total

# Uso de memória
container_memory_usage_bytes
```

### Alertas Recomendados

1. **High Error Rate**: Taxa de erro > 5%
2. **High Latency**: P95 > 1 segundo
3. **Service Down**: Serviço com 0 réplicas
4. **High CPU**: CPU > 80% por 5 minutos
5. **High Memory**: Memória > 90%
6. **Pod Restarts**: Mais de 3 restarts em 10 minutos

---

## ⏮️ Rollback e Disaster Recovery

### Rollback Automático

Ambos os workflows têm rollback automático em caso de falha nos health checks.

### Rollback Manual

#### Docker Swarm

```bash
# Usar script
./deployment/scripts/rollback-swarm.sh

# Ou manualmente
docker service rollback go-ms_order-service
docker service rollback go-ms_payment-service
# ... para cada serviço
```

#### Kubernetes

```bash
# Usar script
./deployment/scripts/rollback-k8s.sh

# Ou manualmente
kubectl rollout undo deployment/order-service -n microservices
kubectl rollout undo deployment/payment-service -n microservices

# Rollback para versão específica
kubectl rollout undo deployment/order-service --to-revision=2 -n microservices

# Ver histórico
kubectl rollout history deployment/order-service -n microservices
```

### Backup e Restore

```bash
# Backup do MySQL
docker exec <mysql-container> mysqldump -u root -p microservices > backup.sql

# Restore
docker exec -i <mysql-container> mysql -u root -p microservices < backup.sql

# Kubernetes
kubectl exec -it mysql-0 -n microservices -- mysqldump -u root -p microservices > backup.sql
kubectl exec -i mysql-0 -n microservices -- mysql -u root -p microservices < backup.sql
```

---

## 🔧 Troubleshooting

### Problemas Comuns

#### 1. Serviço não inicia (Docker Swarm)

```bash
# Ver logs
docker service logs go-ms_order-service

# Ver tasks
docker service ps go-ms_order-service --no-trunc

# Inspecionar serviço
docker service inspect go-ms_order-service
```

#### 2. Pod CrashLoopBackOff (Kubernetes)

```bash
# Ver logs
kubectl logs <pod-name> -n microservices

# Logs anteriores (se restartou)
kubectl logs <pod-name> -n microservices --previous

# Descrever pod
kubectl describe pod <pod-name> -n microservices

# Eventos
kubectl get events -n microservices --sort-by='.lastTimestamp'
```

#### 3. Secret não encontrado

```bash
# Swarm
docker secret ls
docker secret inspect mysql_root_password

# Kubernetes
kubectl get secrets -n microservices
kubectl describe secret mysql-secret -n microservices
```

#### 4. Conexão com DB falha

```bash
# Verificar se MySQL está rodando
# Swarm
docker service ps go-ms_mysql

# Kubernetes
kubectl get pods -l app=mysql -n microservices

# Testar conexão
# Swarm
docker exec -it <mysql-container> mysql -u root -p

# Kubernetes
kubectl exec -it mysql-0 -n microservices -- mysql -u root -p
```

#### 5. Ingress não funciona (Kubernetes)

```bash
# Verificar ingress controller
kubectl get pods -n ingress-nginx

# Ver ingress
kubectl get ingress -n microservices
kubectl describe ingress microservices-ingress -n microservices

# Configurar /etc/hosts
echo "192.168.49.2 microservices.local" | sudo tee -a /etc/hosts
```

### Debug Avançado

```bash
# Entrar em container/pod
# Swarm
docker exec -it <container-id> sh

# Kubernetes
kubectl exec -it <pod-name> -n microservices -- sh

# Ver recursos do cluster
kubectl top nodes
kubectl top pods -n microservices

# Ver uso de recursos (Swarm)
docker stats
```

---

## 📈 Escalabilidade

### Auto-scaling (Kubernetes)

O HPA (Horizontal Pod Autoscaler) já está configurado:

```yaml
# order-service-hpa
minReplicas: 2
maxReplicas: 10
targetCPUUtilizationPercentage: 70
```

### Escalar Manualmente

```bash
# Docker Swarm
docker service scale go-ms_order-service=5

# Kubernetes
kubectl scale deployment order-service --replicas=5 -n microservices
```

---

## ✅ Checklist de Deploy

### Pré-Deploy

- [ ] Todos os testes passaram (CI)
- [ ] Secrets configurados no GitHub
- [ ] Servidor/cluster acessível
- [ ] Backup do banco de dados
- [ ] Tag de versão criada

### Durante Deploy

- [ ] Build de images bem-sucedido
- [ ] Push para registry OK
- [ ] Deploy executado sem erros
- [ ] Health checks passaram
- [ ] Nenhum rollback necessário

### Pós-Deploy

- [ ] Todos os serviços rodando
- [ ] Endpoints acessíveis
- [ ] Métricas sendo coletadas
- [ ] Grafana mostrando dados
- [ ] Jaeger rastreando requests
- [ ] Logs sem erros críticos

---

## 🎯 Próximos Passos

1. **Configurar DNS** para domínios reais
2. **Adicionar SSL/TLS** (Let's Encrypt)
3. **Implementar backup automático**
4. **Configurar alertas** no Grafana
5. **Adicionar rate limiting**
6. **Implementar circuit breaker**
7. **Configurar log aggregation** (ELK, Loki)
8. **Adicionar APM** (Application Performance Monitoring)

---

## 📚 Referências

- [Docker Swarm Documentation](https://docs.docker.com/engine/swarm/)
- [Kubernetes Documentation](https://kubernetes.io/docs/)
- [GitHub Actions](https://docs.github.com/en/actions)
- [Prometheus Best Practices](https://prometheus.io/docs/practices/)
- [12-Factor App](https://12factor.net/)

---

**Última Atualização:** 2024  
**Versão:** 1.0.0  
**Passo:** 12 - Deploy Automatizado em Produção
