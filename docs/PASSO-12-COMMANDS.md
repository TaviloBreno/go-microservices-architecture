# 🔧 Comandos Úteis - Passo 12

## 📑 Índice Rápido

- [Docker Swarm](#-docker-swarm)
- [Kubernetes](#-kubernetes)
- [GitHub Actions](#-github-actions)
- [Monitoramento](#-monitoramento)
- [Debug e Troubleshooting](#-debug-e-troubleshooting)
- [Backup e Restore](#-backup-e-restore)

---

## 🐳 Docker Swarm

### Inicialização

```bash
# Inicializar Swarm
docker swarm init --advertise-addr <IP>

# Obter token para workers
docker swarm join-token worker

# Obter token para managers
docker swarm join-token manager

# Adicionar worker (executar no node worker)
docker swarm join --token <TOKEN> <MANAGER-IP>:2377
```

### Gerenciamento de Nodes

```bash
# Listar nodes
docker node ls

# Inspecionar node
docker node inspect <NODE-ID>

# Promover worker a manager
docker node promote <NODE-ID>

# Rebaixar manager a worker
docker node demote <NODE-ID>

# Drenar node (mover containers para outros nodes)
docker node update --availability drain <NODE-ID>

# Ativar node novamente
docker node update --availability active <NODE-ID>

# Remover node
docker node rm <NODE-ID>
```

### Deploy e Atualização

```bash
# Deploy do stack
docker stack deploy -c deployment/docker-swarm/stack.yml --with-registry-auth go-ms

# Listar stacks
docker stack ls

# Listar serviços do stack
docker stack services go-ms

# Ver tasks (containers) do stack
docker stack ps go-ms

# Remover stack
docker stack rm go-ms

# Atualizar imagem de um serviço
docker service update --image tavilobreno/go-ms-order:v1.1.0 go-ms_order-service

# Escalar serviço
docker service scale go-ms_order-service=5

# Forçar update (recriar containers)
docker service update --force go-ms_order-service

# Rollback
docker service rollback go-ms_order-service
```

### Logs e Debug

```bash
# Logs de um serviço
docker service logs go-ms_order-service

# Seguir logs (tail -f)
docker service logs -f go-ms_order-service

# Últimas 100 linhas
docker service logs --tail 100 go-ms_order-service

# Logs com timestamps
docker service logs -t go-ms_order-service

# Inspecionar serviço
docker service inspect go-ms_order-service

# Ver tasks com detalhes
docker service ps go-ms_order-service --no-trunc

# Ver apenas tasks rodando
docker service ps go-ms_order-service --filter "desired-state=running"

# Ver containers em todos os nodes
docker ps --filter "label=com.docker.swarm.service.name=go-ms_order-service"
```

### Secrets

```bash
# Criar secret
echo "senha123" | docker secret create mysql_root_password -

# Criar secret de arquivo
docker secret create mysql_config ./mysql.cnf

# Listar secrets
docker secret ls

# Inspecionar secret (não mostra valor)
docker secret inspect mysql_root_password

# Remover secret
docker secret rm mysql_root_password
```

### Networks

```bash
# Listar networks
docker network ls

# Inspecionar network overlay
docker network inspect go-ms_microservices

# Ver containers em uma network
docker network inspect go-ms_microservices -f '{{range .Containers}}{{.Name}} {{end}}'
```

### Volumes

```bash
# Listar volumes
docker volume ls

# Inspecionar volume
docker volume inspect go-ms_mysql-data

# Remover volume
docker volume rm go-ms_mysql-data

# Remover volumes não usados
docker volume prune
```

---

## ☸️ Kubernetes

### Cluster

```bash
# Verificar conexão
kubectl cluster-info

# Ver nodes
kubectl get nodes

# Descrever node
kubectl describe node <NODE-NAME>

# Ver recursos do cluster
kubectl top nodes
kubectl top pods -n microservices

# Contextos (clusters)
kubectl config get-contexts
kubectl config use-context <CONTEXT-NAME>

# Namespace atual
kubectl config view --minify --output 'jsonpath={..namespace}'

# Definir namespace padrão
kubectl config set-context --current --namespace=microservices
```

### Deploy e Atualização

```bash
# Apply de todos os manifests
kubectl apply -f deployment/kubernetes/

# Apply de um arquivo específico
kubectl apply -f deployment/kubernetes/10-order-service.yaml

# Delete
kubectl delete -f deployment/kubernetes/10-order-service.yaml

# Namespace
kubectl create namespace microservices
kubectl delete namespace microservices

# Ver recursos em um namespace
kubectl get all -n microservices

# Ver recursos em todos os namespaces
kubectl get all --all-namespaces
```

### Deployments

```bash
# Listar deployments
kubectl get deployments -n microservices

# Descrever deployment
kubectl describe deployment order-service -n microservices

# Escalar deployment
kubectl scale deployment order-service --replicas=5 -n microservices

# Atualizar imagem
kubectl set image deployment/order-service \
  order-service=tavilobreno/go-ms-order:v1.1.0 -n microservices

# Editar deployment
kubectl edit deployment order-service -n microservices

# Delete deployment
kubectl delete deployment order-service -n microservices
```

### Rollout

```bash
# Status do rollout
kubectl rollout status deployment/order-service -n microservices

# Histórico de rollout
kubectl rollout history deployment/order-service -n microservices

# Detalhes de uma revisão
kubectl rollout history deployment/order-service --revision=2 -n microservices

# Rollback para versão anterior
kubectl rollout undo deployment/order-service -n microservices

# Rollback para revisão específica
kubectl rollout undo deployment/order-service --to-revision=2 -n microservices

# Pausar rollout
kubectl rollout pause deployment/order-service -n microservices

# Continuar rollout
kubectl rollout resume deployment/order-service -n microservices

# Restart deployment (força recriação de pods)
kubectl rollout restart deployment/order-service -n microservices
```

### Pods

```bash
# Listar pods
kubectl get pods -n microservices

# Pods com mais detalhes
kubectl get pods -n microservices -o wide

# Descrever pod
kubectl describe pod <POD-NAME> -n microservices

# Logs de um pod
kubectl logs <POD-NAME> -n microservices

# Logs seguindo (tail -f)
kubectl logs -f <POD-NAME> -n microservices

# Logs do container anterior (se restartou)
kubectl logs <POD-NAME> -n microservices --previous

# Logs de múltiplos pods (por label)
kubectl logs -l app=order-service -n microservices

# Exec em pod
kubectl exec -it <POD-NAME> -n microservices -- sh
kubectl exec -it <POD-NAME> -n microservices -- /bin/bash

# Ver uso de recursos
kubectl top pod <POD-NAME> -n microservices

# Delete pod (recriado automaticamente se parte de deployment)
kubectl delete pod <POD-NAME> -n microservices

# Forçar delete (quando fica stuck)
kubectl delete pod <POD-NAME> -n microservices --force --grace-period=0
```

### Services

```bash
# Listar services
kubectl get services -n microservices

# Descrever service
kubectl describe service order-service -n microservices

# Ver endpoints do service
kubectl get endpoints order-service -n microservices

# Port-forward para acessar localmente
kubectl port-forward service/order-service 8080:8080 -n microservices

# Edit service
kubectl edit service order-service -n microservices
```

### Secrets

```bash
# Criar secret genérico
kubectl create secret generic mysql-secret -n microservices \
  --from-literal=root-password=root123 \
  --from-literal=user=microservices \
  --from-literal=password=micro123 \
  --from-literal=database=microservices

# Criar secret de arquivo
kubectl create secret generic tls-secret -n microservices \
  --from-file=tls.crt=./tls.crt \
  --from-file=tls.key=./tls.key

# Criar secret TLS
kubectl create secret tls tls-secret -n microservices \
  --cert=./tls.crt \
  --key=./tls.key

# Listar secrets
kubectl get secrets -n microservices

# Descrever secret
kubectl describe secret mysql-secret -n microservices

# Ver secret em base64
kubectl get secret mysql-secret -n microservices -o yaml

# Decodificar valor de secret
kubectl get secret mysql-secret -n microservices -o jsonpath='{.data.root-password}' | base64 -d

# Delete secret
kubectl delete secret mysql-secret -n microservices
```

### ConfigMaps

```bash
# Criar configmap
kubectl create configmap app-config -n microservices \
  --from-literal=ENV=production \
  --from-literal=DEBUG=false

# Criar configmap de arquivo
kubectl create configmap prometheus-config -n microservices \
  --from-file=prometheus.yml

# Listar configmaps
kubectl get configmaps -n microservices

# Ver conteúdo
kubectl describe configmap app-config -n microservices
kubectl get configmap app-config -n microservices -o yaml

# Edit configmap
kubectl edit configmap app-config -n microservices

# Delete configmap
kubectl delete configmap app-config -n microservices
```

### StatefulSets

```bash
# Listar statefulsets
kubectl get statefulsets -n microservices

# Descrever statefulset
kubectl describe statefulset mysql -n microservices

# Escalar statefulset
kubectl scale statefulset mysql --replicas=3 -n microservices

# Delete statefulset
kubectl delete statefulset mysql -n microservices

# Delete mas manter pods
kubectl delete statefulset mysql -n microservices --cascade=orphan
```

### PersistentVolumeClaims (PVCs)

```bash
# Listar PVCs
kubectl get pvc -n microservices

# Descrever PVC
kubectl describe pvc mysql-data-mysql-0 -n microservices

# Ver PVs (persistent volumes)
kubectl get pv

# Delete PVC
kubectl delete pvc mysql-data-mysql-0 -n microservices
```

### Ingress

```bash
# Listar ingress
kubectl get ingress -n microservices

# Descrever ingress
kubectl describe ingress microservices-ingress -n microservices

# Edit ingress
kubectl edit ingress microservices-ingress -n microservices

# Ver IP/hostname do ingress
kubectl get ingress microservices-ingress -n microservices \
  -o jsonpath='{.status.loadBalancer.ingress[0].ip}'
```

### HPA (Horizontal Pod Autoscaler)

```bash
# Listar HPAs
kubectl get hpa -n microservices

# Descrever HPA
kubectl describe hpa order-service-hpa -n microservices

# Criar HPA
kubectl autoscale deployment order-service -n microservices \
  --cpu-percent=70 --min=2 --max=10

# Edit HPA
kubectl edit hpa order-service-hpa -n microservices

# Delete HPA
kubectl delete hpa order-service-hpa -n microservices
```

### Eventos

```bash
# Ver eventos recentes
kubectl get events -n microservices

# Eventos ordenados por tempo
kubectl get events -n microservices --sort-by='.lastTimestamp'

# Eventos de um pod específico
kubectl get events -n microservices --field-selector involvedObject.name=<POD-NAME>
```

---

## 🔄 GitHub Actions

### Comandos Locais

```bash
# Instalar act (rodar workflows localmente)
# macOS
brew install act

# Linux
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Listar workflows
act -l

# Rodar workflow
act -j deploy-swarm

# Rodar com secrets
act -j deploy-swarm --secret-file .secrets
```

### Git Tags

```bash
# Criar tag
git tag v1.0.0

# Criar tag anotada
git tag -a v1.0.0 -m "Release version 1.0.0"

# Listar tags
git tag
git tag -l "v1.*"

# Ver detalhes da tag
git show v1.0.0

# Push de tag (dispara deploy)
git push origin v1.0.0

# Push de todas as tags
git push --tags

# Delete tag local
git tag -d v1.0.0

# Delete tag remota
git push origin :refs/tags/v1.0.0
```

### Workflow Dispatch (Manual)

```bash
# Via GitHub CLI
gh workflow run deploy-swarm.yml

# Com inputs
gh workflow run deploy-kubernetes.yml -f environment=production

# Ver runs
gh run list --workflow=deploy-swarm.yml

# Ver logs de um run
gh run view <RUN-ID> --log

# Cancelar run
gh run cancel <RUN-ID>
```

---

## 📊 Monitoramento

### Prometheus

```bash
# Queries úteis via API
# Taxa de requisições
curl 'http://localhost:9090/api/v1/query?query=rate(grpc_server_handled_total[5m])'

# Latência P95
curl 'http://localhost:9090/api/v1/query?query=histogram_quantile(0.95,rate(grpc_server_handling_seconds_bucket[5m]))'

# Status do Prometheus
curl http://localhost:9090/-/healthy
curl http://localhost:9090/-/ready

# Reload config
curl -X POST http://localhost:9090/-/reload
```

### Grafana

```bash
# Health check
curl http://localhost:3000/api/health

# Login (obter token)
curl -X POST http://localhost:3000/api/auth/keys \
  -H "Content-Type: application/json" \
  -u admin:admin123 \
  -d '{"name":"apikey","role":"Admin"}'

# Listar datasources
curl http://localhost:3000/api/datasources \
  -H "Authorization: Bearer <TOKEN>"

# Listar dashboards
curl http://localhost:3000/api/search \
  -H "Authorization: Bearer <TOKEN>"
```

### Jaeger

```bash
# Health check
curl http://localhost:16686

# Buscar traces via API
curl 'http://localhost:16686/api/traces?service=order-service&limit=10'

# Buscar serviços
curl http://localhost:16686/api/services
```

### RabbitMQ

```bash
# API Management
# Listar queues
curl -u guest:guest123 http://localhost:15672/api/queues

# Listar exchanges
curl -u guest:guest123 http://localhost:15672/api/exchanges

# Listar conexões
curl -u guest:guest123 http://localhost:15672/api/connections

# Health check
curl -u guest:guest123 http://localhost:15672/api/health/checks/alarms
```

---

## 🐛 Debug e Troubleshooting

### Docker Swarm

```bash
# Ver porque serviço não inicia
docker service ps go-ms_order-service --no-trunc

# Inspecionar task que falhou
docker inspect <TASK-ID>

# Ver todos os containers (incluindo stopped)
docker ps -a --filter "label=com.docker.swarm.service.name=go-ms_order-service"

# Logs de container específico
docker logs <CONTAINER-ID>

# Exec em container
docker exec -it <CONTAINER-ID> sh

# Ver uso de recursos em tempo real
docker stats

# Ver configuração do stack
docker stack config -c deployment/docker-swarm/stack.yml

# Validar stack.yml
docker-compose -f deployment/docker-swarm/stack.yml config
```

### Kubernetes

```bash
# Ver porque pod não inicia
kubectl describe pod <POD-NAME> -n microservices

# Ver eventos de erro
kubectl get events -n microservices | grep Error
kubectl get events -n microservices | grep Warning

# Ver logs de init container
kubectl logs <POD-NAME> -c <INIT-CONTAINER-NAME> -n microservices

# Debug de pod (cria pod temporário)
kubectl run debug --image=busybox -it --rm --restart=Never -n microservices -- sh

# Testar conectividade entre pods
kubectl run curl --image=curlimages/curl -it --rm --restart=Never -n microservices \
  -- curl http://order-service:8080/health

# Verificar DNS
kubectl run dnsutils --image=gcr.io/kubernetes-e2e-test-images/dnsutils:1.3 \
  -it --rm --restart=Never -n microservices -- nslookup order-service

# Ver configuração efetiva de um pod
kubectl get pod <POD-NAME> -n microservices -o yaml

# Ver limites de recursos
kubectl describe nodes | grep -A 5 "Allocated resources"

# Ver políticas de rede
kubectl get networkpolicies -n microservices
kubectl describe networkpolicy microservices-network-policy -n microservices
```

### Conectividade

```bash
# Testar endpoint de dentro do cluster
# Swarm
docker run --rm --network go-ms_microservices curlimages/curl \
  curl http://order-service:8080/health

# Kubernetes
kubectl run curl --image=curlimages/curl -it --rm --restart=Never -n microservices \
  -- curl http://order-service:8080/health

# Port-forward para debug local (K8s)
kubectl port-forward service/order-service 8080:8080 -n microservices
# Agora acesse: http://localhost:8080/health

# Port-forward de pod
kubectl port-forward <POD-NAME> 8080:8080 -n microservices
```

### Database

```bash
# MySQL
# Swarm
docker exec -it <mysql-container> mysql -u root -p

# Kubernetes
kubectl exec -it mysql-0 -n microservices -- mysql -u root -p

# Queries úteis
SHOW DATABASES;
USE microservices;
SHOW TABLES;
SELECT * FROM orders LIMIT 10;

# Verificar conexões
SHOW PROCESSLIST;

# Ver status
SHOW STATUS;
```

---

## 💾 Backup e Restore

### MySQL

```bash
# Backup
# Swarm
docker exec <mysql-container> mysqldump -u root -p microservices > backup.sql

# Kubernetes
kubectl exec mysql-0 -n microservices -- mysqldump -u root -p microservices > backup.sql

# Backup com compressão
kubectl exec mysql-0 -n microservices -- mysqldump -u root -p microservices | gzip > backup.sql.gz

# Restore
# Swarm
docker exec -i <mysql-container> mysql -u root -p microservices < backup.sql

# Kubernetes
kubectl exec -i mysql-0 -n microservices -- mysql -u root -p microservices < backup.sql

# Restore com gunzip
gunzip < backup.sql.gz | kubectl exec -i mysql-0 -n microservices -- mysql -u root -p microservices
```

### Volumes

```bash
# Backup de volume (Docker)
docker run --rm \
  -v go-ms_mysql-data:/source:ro \
  -v $(pwd):/backup \
  alpine tar czf /backup/mysql-backup.tar.gz -C /source .

# Restore de volume
docker run --rm \
  -v go-ms_mysql-data:/target \
  -v $(pwd):/backup \
  alpine tar xzf /backup/mysql-backup.tar.gz -C /target

# Kubernetes PVC backup (usando pod temporário)
kubectl run backup --image=alpine -n microservices \
  --overrides='
  {
    "spec": {
      "containers": [{
        "name": "backup",
        "image": "alpine",
        "command": ["tar", "czf", "/backup/mysql-backup.tar.gz", "-C", "/data", "."],
        "volumeMounts": [
          {"name": "data", "mountPath": "/data"},
          {"name": "backup", "mountPath": "/backup"}
        ]
      }],
      "volumes": [
        {"name": "data", "persistentVolumeClaim": {"claimName": "mysql-data-mysql-0"}},
        {"name": "backup", "hostPath": {"path": "/tmp/backup"}}
      ],
      "restartPolicy": "Never"
    }
  }'
```

### Configurações

```bash
# Backup de secrets (K8s)
kubectl get secrets -n microservices -o yaml > secrets-backup.yaml

# Backup de configmaps
kubectl get configmaps -n microservices -o yaml > configmaps-backup.yaml

# Backup de deployments
kubectl get deployments -n microservices -o yaml > deployments-backup.yaml

# Backup completo do namespace
kubectl get all,secrets,configmaps,pvc -n microservices -o yaml > namespace-backup.yaml

# Restore
kubectl apply -f namespace-backup.yaml
```

---

## 🧪 Scripts Úteis

### Health Check Rápido

```bash
# Docker Swarm
./deployment/scripts/health-check-swarm.sh

# Kubernetes
./deployment/scripts/health-check-k8s.sh

# Manual (Swarm)
curl http://localhost:8080/health
curl http://localhost:9090/-/healthy
curl http://localhost:3000/api/health

# Manual (K8s)
kubectl run curl --image=curlimages/curl -it --rm --restart=Never -n microservices \
  -- curl http://order-service:8080/health
```

### Deploy Rápido

```bash
# Swarm
./deployment/scripts/deploy-swarm.sh

# Kubernetes
./deployment/scripts/deploy-k8s.sh

# Via GitHub (criar tag)
git tag v1.0.0 && git push origin v1.0.0
```

### Rollback Rápido

```bash
# Swarm
./deployment/scripts/rollback-swarm.sh

# Kubernetes
./deployment/scripts/rollback-k8s.sh
```

### Cleanup Completo

```bash
# Docker Swarm
docker stack rm go-ms
docker secret rm $(docker secret ls -q)
docker volume prune -f
docker network prune -f

# Kubernetes
kubectl delete namespace microservices
kubectl delete pv --all
```

---

## 📝 Cheat Sheet

### Top 10 Comandos Mais Usados

```bash
# 1. Ver status geral
docker stack services go-ms                    # Swarm
kubectl get all -n microservices              # K8s

# 2. Ver logs
docker service logs -f go-ms_order-service    # Swarm
kubectl logs -f <pod> -n microservices        # K8s

# 3. Escalar serviço
docker service scale go-ms_order-service=5    # Swarm
kubectl scale deployment order-service --replicas=5 -n microservices  # K8s

# 4. Update de imagem
docker service update --image tavilobreno/go-ms-order:v1.1.0 go-ms_order-service  # Swarm
kubectl set image deployment/order-service order-service=tavilobreno/go-ms-order:v1.1.0 -n microservices  # K8s

# 5. Rollback
docker service rollback go-ms_order-service                      # Swarm
kubectl rollout undo deployment/order-service -n microservices  # K8s

# 6. Health check
curl http://localhost:8080/health                    # Swarm
kubectl run curl --image=curlimages/curl -it --rm --restart=Never -n microservices -- curl http://order-service:8080/health  # K8s

# 7. Exec em container/pod
docker exec -it <container-id> sh                              # Swarm
kubectl exec -it <pod-name> -n microservices -- sh            # K8s

# 8. Ver eventos
docker service ps go-ms_order-service --no-trunc               # Swarm
kubectl get events -n microservices --sort-by='.lastTimestamp'  # K8s

# 9. Port-forward (debug)
# N/A para Swarm (usar portas publicadas)
kubectl port-forward service/order-service 8080:8080 -n microservices  # K8s

# 10. Deploy/Redeploy
docker stack deploy -c deployment/docker-swarm/stack.yml --with-registry-auth go-ms  # Swarm
kubectl apply -f deployment/kubernetes/ -n microservices                             # K8s
```

---

**Última Atualização:** 2024  
**Versão:** 1.0.0  
**Passo:** 12 - Deploy Automatizado em Produção
