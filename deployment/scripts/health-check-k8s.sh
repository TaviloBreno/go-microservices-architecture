#!/bin/bash

# ========================================
# Kubernetes Health Check Script
# ========================================

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

NAMESPACE="${NAMESPACE:-microservices}"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Kubernetes Health Check${NC}"
echo -e "${BLUE}========================================${NC}\n"

total_checks=0
passed_checks=0
failed_checks=0

# Check cluster connection
echo -e "${YELLOW}Checking cluster connection...${NC}"
if kubectl cluster-info &>/dev/null; then
    echo -e "${GREEN}✓${NC} Connected to cluster\n"
else
    echo -e "${RED}❌ Cannot connect to cluster${NC}"
    exit 1
fi

# Check namespace
echo -e "${YELLOW}Checking namespace...${NC}"
if kubectl get namespace $NAMESPACE &>/dev/null; then
    echo -e "${GREEN}✓${NC} Namespace exists: $NAMESPACE\n"
else
    echo -e "${RED}❌ Namespace not found: $NAMESPACE${NC}"
    exit 1
fi

# Check Deployments
echo -e "${YELLOW}Checking Deployments:${NC}\n"

DEPLOYMENTS=$(kubectl get deployments -n $NAMESPACE -o jsonpath='{.items[*].metadata.name}')

for DEPLOYMENT in $DEPLOYMENTS; do
    ((total_checks++))
    
    READY=$(kubectl get deployment $DEPLOYMENT -n $NAMESPACE -o jsonpath='{.status.readyReplicas}')
    DESIRED=$(kubectl get deployment $DEPLOYMENT -n $NAMESPACE -o jsonpath='{.status.replicas}')
    
    if [ "$READY" = "$DESIRED" ] && [ "$READY" != "" ]; then
        echo -e "${GREEN}✓${NC} $DEPLOYMENT: $READY/$DESIRED replicas ready"
        ((passed_checks++))
    else
        echo -e "${RED}✗${NC} $DEPLOYMENT: $READY/$DESIRED replicas ready"
        ((failed_checks++))
    fi
done

# Check Pods
echo -e "\n${YELLOW}Checking Pods:${NC}\n"

UNHEALTHY_PODS=$(kubectl get pods -n $NAMESPACE --field-selector=status.phase!=Running -o name 2>/dev/null)

if [ -z "$UNHEALTHY_PODS" ]; then
    echo -e "${GREEN}✓${NC} All pods are running"
    ((passed_checks++))
else
    echo -e "${RED}✗${NC} Unhealthy pods detected:"
    echo "$UNHEALTHY_PODS"
    ((failed_checks++))
    
    # Show pod details
    for POD in $UNHEALTHY_PODS; do
        echo -e "\n${YELLOW}Details for $POD:${NC}"
        kubectl describe $POD -n $NAMESPACE | tail -20
    done
fi

((total_checks++))

# Check Services
echo -e "\n${YELLOW}Checking Services:${NC}\n"

SERVICES=$(kubectl get services -n $NAMESPACE -o jsonpath='{.items[*].metadata.name}')

for SERVICE in $SERVICES; do
    ((total_checks++))
    
    ENDPOINTS=$(kubectl get endpoints $SERVICE -n $NAMESPACE -o jsonpath='{.subsets[*].addresses[*].ip}' 2>/dev/null)
    
    if [ -n "$ENDPOINTS" ]; then
        COUNT=$(echo $ENDPOINTS | wc -w)
        echo -e "${GREEN}✓${NC} $SERVICE: $COUNT endpoint(s)"
        ((passed_checks++))
    else
        echo -e "${YELLOW}⚠${NC} $SERVICE: No endpoints"
        # Don't count as failure for headless services
    fi
done

# Check Ingress
echo -e "\n${YELLOW}Checking Ingress:${NC}\n"

INGRESS_COUNT=$(kubectl get ingress -n $NAMESPACE --no-headers 2>/dev/null | wc -l)

if [ "$INGRESS_COUNT" -gt 0 ]; then
    echo -e "${GREEN}✓${NC} Ingress configured ($INGRESS_COUNT)"
    kubectl get ingress -n $NAMESPACE
else
    echo -e "${YELLOW}⚠${NC} No ingress configured"
fi

# Check PVCs
echo -e "\n${YELLOW}Checking Persistent Volume Claims:${NC}\n"

PVC_STATUS=$(kubectl get pvc -n $NAMESPACE --no-headers 2>/dev/null)

if [ -n "$PVC_STATUS" ]; then
    echo "$PVC_STATUS"
    
    UNBOUND=$(echo "$PVC_STATUS" | grep -v Bound || true)
    if [ -n "$UNBOUND" ]; then
        echo -e "${RED}✗${NC} Some PVCs are not bound"
        ((failed_checks++))
    else
        echo -e "${GREEN}✓${NC} All PVCs are bound"
        ((passed_checks++))
    fi
    ((total_checks++))
fi

# Summary
echo -e "\n${BLUE}========================================${NC}"
echo -e "${BLUE}  Health Check Summary${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "Total Checks:   $total_checks"
echo -e "${GREEN}Passed:${NC}         $passed_checks"
echo -e "${RED}Failed:${NC}         $failed_checks"
echo ""

if [ $total_checks -eq 0 ]; then
    echo -e "${YELLOW}⚠ No checks performed${NC}"
    exit 1
fi

success_rate=$(( (passed_checks * 100) / total_checks ))

if [ $success_rate -eq 100 ]; then
    echo -e "${GREEN}🎉 All checks passed! (100%)${NC}"
    exit 0
elif [ $success_rate -ge 80 ]; then
    echo -e "${YELLOW}⚠ Most checks passed ($success_rate%)${NC}"
    exit 0
else
    echo -e "${RED}❌ Many checks failed ($success_rate%)${NC}"
    exit 1
fi
