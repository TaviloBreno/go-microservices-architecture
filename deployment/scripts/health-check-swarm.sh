#!/bin/bash

# ========================================
# Docker Swarm Health Check Script
# ========================================

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

STACK_NAME="${STACK_NAME:-go-ms}"
HOST="${SWARM_HOST:-localhost}"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Docker Swarm Health Check${NC}"
echo -e "${BLUE}========================================${NC}\n"

total_checks=0
passed_checks=0
failed_checks=0

# Function to check HTTP endpoint
check_http() {
    local url=$1
    local service=$2
    
    ((total_checks++))
    
    if curl -sf "$url" > /dev/null 2>&1; then
        echo -e "${GREEN}✓${NC} $service is ${GREEN}healthy${NC}"
        ((passed_checks++))
        return 0
    else
        echo -e "${RED}✗${NC} $service is ${RED}unhealthy${NC}"
        ((failed_checks++))
        return 1
    fi
}

# Check Docker Stack Services
echo -e "${YELLOW}Checking Docker Stack Services:${NC}\n"
docker stack services $STACK_NAME

echo -e "\n${YELLOW}Checking Service Health:${NC}\n"

# Check microservices
check_http "http://$HOST:8080/health" "BFF GraphQL"
check_http "http://$HOST:3001" "Frontend"

# Check monitoring
check_http "http://$HOST:9090/-/healthy" "Prometheus"
check_http "http://$HOST:3000/api/health" "Grafana"
check_http "http://$HOST:16686" "Jaeger UI"

# Check infrastructure
check_http "http://$HOST:15672" "RabbitMQ Management"

# Check for service replicas
echo -e "\n${YELLOW}Checking Service Replicas:${NC}\n"

FAILED_SERVICES=$(docker stack services $STACK_NAME --format '{{.Name}}: {{.Replicas}}' | grep '0/' || true)

if [ -n "$FAILED_SERVICES" ]; then
    echo -e "${RED}⚠ Services with no replicas:${NC}"
    echo "$FAILED_SERVICES"
    ((failed_checks++))
else
    echo -e "${GREEN}✓${NC} All services have running replicas"
    ((passed_checks++))
fi

# Summary
echo -e "\n${BLUE}========================================${NC}"
echo -e "${BLUE}  Health Check Summary${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "Total Checks:   $total_checks"
echo -e "${GREEN}Passed:${NC}         $passed_checks"
echo -e "${RED}Failed:${NC}         $failed_checks"
echo ""

success_rate=$(( (passed_checks * 100) / total_checks ))

if [ $success_rate -eq 100 ]; then
    echo -e "${GREEN}🎉 All services are healthy! (100%)${NC}"
    exit 0
elif [ $success_rate -ge 80 ]; then
    echo -e "${YELLOW}⚠ Most services are healthy ($success_rate%)${NC}"
    exit 0
else
    echo -e "${RED}❌ Multiple services are down ($success_rate%)${NC}"
    exit 1
fi
