#!/bin/bash

# ========================================
# Kubernetes Rollback Script
# ========================================

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

NAMESPACE="${NAMESPACE:-microservices}"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Kubernetes Rollback${NC}"
echo -e "${BLUE}========================================${NC}\n"

echo -e "${RED}⚠️  WARNING: This will rollback all deployments to their previous version${NC}"
read -p "Are you sure you want to continue? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo -e "${YELLOW}Rollback cancelled${NC}"
    exit 0
fi

echo -e "\n${YELLOW}Rolling back deployments...${NC}\n"

DEPLOYMENTS=(
    "order-service"
    "payment-service"
    "user-service"
    "notification-service"
    "catalog-service"
    "bff-service"
    "frontend"
)

for DEPLOYMENT in "${DEPLOYMENTS[@]}"; do
    echo -e "${BLUE}Rolling back: $DEPLOYMENT${NC}"
    
    if kubectl rollout undo deployment/$DEPLOYMENT -n $NAMESPACE; then
        echo -e "${GREEN}✓${NC} Rolled back: $DEPLOYMENT"
    else
        echo -e "${RED}❌ Failed to rollback: $DEPLOYMENT${NC}"
    fi
    
    sleep 2
done

echo -e "\n${YELLOW}Waiting for rollout to complete...${NC}\n"

for DEPLOYMENT in "${DEPLOYMENTS[@]}"; do
    echo -e "${BLUE}Checking: $DEPLOYMENT${NC}"
    kubectl rollout status deployment/$DEPLOYMENT -n $NAMESPACE --timeout=120s || true
done

# Show deployment status
echo -e "\n${BLUE}Deployment Status After Rollback:${NC}"
kubectl get deployments -n $NAMESPACE

echo -e "\n${BLUE}Pod Status:${NC}"
kubectl get pods -n $NAMESPACE

echo -e "\n${GREEN}Rollback completed!${NC}"
echo -e "${YELLOW}Check pod logs if any issues persist${NC}"
