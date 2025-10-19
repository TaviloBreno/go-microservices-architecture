#!/bin/bash

# ========================================
# Kubernetes Deployment Script
# ========================================

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

NAMESPACE="${NAMESPACE:-microservices}"
VERSION="${VERSION:-latest}"
IMAGE_REGISTRY="${IMAGE_REGISTRY:-tavilobreno}"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Kubernetes Deployment${NC}"
echo -e "${BLUE}========================================${NC}\n"

# Check kubectl
if ! command -v kubectl &> /dev/null; then
    echo -e "${RED}❌ kubectl not found${NC}"
    exit 1
fi

# Check cluster connection
echo -e "${YELLOW}Checking cluster connection...${NC}"
if ! kubectl cluster-info &>/dev/null; then
    echo -e "${RED}❌ Cannot connect to Kubernetes cluster${NC}"
    exit 1
fi

echo -e "${GREEN}✓${NC} Connected to cluster\n"

# Create namespace
echo -e "${YELLOW}Creating namespace...${NC}"
kubectl apply -f deployment/kubernetes/00-namespace.yaml
echo -e "${GREEN}✓${NC} Namespace created/verified\n"

# Create secrets
echo -e "${YELLOW}Setting up secrets...${NC}"

if [ -z "$MYSQL_ROOT_PASSWORD" ]; then
    read -sp "Enter MySQL root password: " MYSQL_ROOT_PASSWORD
    echo
fi

if [ -z "$MYSQL_PASSWORD" ]; then
    read -sp "Enter MySQL user password: " MYSQL_PASSWORD
    echo
fi

if [ -z "$RABBITMQ_PASSWORD" ]; then
    read -sp "Enter RabbitMQ password: " RABBITMQ_PASSWORD
    echo
fi

if [ -z "$GRAFANA_PASSWORD" ]; then
    read -sp "Enter Grafana admin password: " GRAFANA_PASSWORD
    echo
fi

if [ -z "$SMTP_PASSWORD" ]; then
    read -sp "Enter SMTP password (optional): " SMTP_PASSWORD
    echo
fi

# Delete existing secrets
kubectl delete secret mysql-secret -n $NAMESPACE --ignore-not-found=true
kubectl delete secret rabbitmq-secret -n $NAMESPACE --ignore-not-found=true
kubectl delete secret grafana-secret -n $NAMESPACE --ignore-not-found=true
kubectl delete secret smtp-secret -n $NAMESPACE --ignore-not-found=true

# Create secrets
kubectl create secret generic mysql-secret -n $NAMESPACE \
    --from-literal=root-password="$MYSQL_ROOT_PASSWORD" \
    --from-literal=user=microservices \
    --from-literal=password="$MYSQL_PASSWORD" \
    --from-literal=database=microservices

kubectl create secret generic rabbitmq-secret -n $NAMESPACE \
    --from-literal=user=guest \
    --from-literal=password="$RABBITMQ_PASSWORD" \
    --from-literal=url="amqp://guest:${RABBITMQ_PASSWORD}@rabbitmq:5672/"

kubectl create secret generic grafana-secret -n $NAMESPACE \
    --from-literal=admin-user=admin \
    --from-literal=admin-password="$GRAFANA_PASSWORD"

[ -n "$SMTP_PASSWORD" ] && kubectl create secret generic smtp-secret -n $NAMESPACE \
    --from-literal=host=smtp.gmail.com \
    --from-literal=port=587 \
    --from-literal=password="$SMTP_PASSWORD"

echo -e "${GREEN}✓${NC} Secrets created\n"

# Apply ConfigMaps
echo -e "${YELLOW}Applying ConfigMaps...${NC}"
kubectl apply -f deployment/kubernetes/02-configmaps.yaml
echo -e "${GREEN}✓${NC} ConfigMaps applied\n"

# Deploy Infrastructure
echo -e "${YELLOW}Deploying infrastructure (MySQL, RabbitMQ)...${NC}"
kubectl apply -f deployment/kubernetes/30-infrastructure.yaml
echo -e "${BLUE}⏳ Waiting for infrastructure to be ready...${NC}"
sleep 30
echo -e "${GREEN}✓${NC} Infrastructure deployed\n"

# Deploy Microservices
echo -e "${YELLOW}Deploying microservices...${NC}"
kubectl apply -f deployment/kubernetes/10-order-service.yaml
kubectl apply -f deployment/kubernetes/11-payment-service.yaml
kubectl apply -f deployment/kubernetes/12-other-services.yaml
echo -e "${BLUE}⏳ Waiting for microservices...${NC}"
sleep 20
echo -e "${GREEN}✓${NC} Microservices deployed\n"

# Deploy BFF and Frontend
echo -e "${YELLOW}Deploying BFF and Frontend...${NC}"
kubectl apply -f deployment/kubernetes/20-bff-frontend.yaml
echo -e "${GREEN}✓${NC} BFF and Frontend deployed\n"

# Deploy Monitoring
echo -e "${YELLOW}Deploying monitoring stack...${NC}"
kubectl apply -f deployment/kubernetes/40-monitoring.yaml
echo -e "${GREEN}✓${NC} Monitoring deployed\n"

# Deploy Ingress and HPA
echo -e "${YELLOW}Deploying Ingress and Autoscaling...${NC}"
kubectl apply -f deployment/kubernetes/50-ingress.yaml
echo -e "${GREEN}✓${NC} Ingress and HPA deployed\n"

# Wait for rollout
echo -e "${YELLOW}Waiting for deployments to complete...${NC}"
kubectl rollout status deployment/order-service -n $NAMESPACE --timeout=300s || true
kubectl rollout status deployment/payment-service -n $NAMESPACE --timeout=300s || true
kubectl rollout status deployment/user-service -n $NAMESPACE --timeout=300s || true
kubectl rollout status deployment/bff-service -n $NAMESPACE --timeout=300s || true
kubectl rollout status deployment/frontend -n $NAMESPACE --timeout=300s || true

echo ""

# Show status
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Deployment Status${NC}"
echo -e "${BLUE}========================================${NC}\n"

echo -e "${YELLOW}Deployments:${NC}"
kubectl get deployments -n $NAMESPACE

echo -e "\n${YELLOW}Pods:${NC}"
kubectl get pods -n $NAMESPACE

echo -e "\n${YELLOW}Services:${NC}"
kubectl get services -n $NAMESPACE

echo -e "\n${YELLOW}Ingress:${NC}"
kubectl get ingress -n $NAMESPACE

# Show endpoints
echo -e "\n${BLUE}========================================${NC}"
echo -e "${BLUE}  Access URLs${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "${YELLOW}Configure /etc/hosts with your cluster IP:${NC}"
echo -e "${GREEN}Frontend:${NC}        http://microservices.local"
echo -e "${GREEN}GraphQL API:${NC}     http://api.microservices.local"
echo -e "${GREEN}Grafana:${NC}         http://grafana.microservices.local"
echo -e "${GREEN}Prometheus:${NC}      http://prometheus.microservices.local"
echo -e "${GREEN}Jaeger:${NC}          http://jaeger.microservices.local"
echo ""

echo -e "${GREEN}🎉 Deployment completed!${NC}"
echo -e "${YELLOW}Run './deployment/scripts/health-check-k8s.sh' to verify all services${NC}"
