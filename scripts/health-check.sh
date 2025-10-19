#!/bin/bash

# Health Check Script for Microservices
# Verifica a saúde de todos os serviços e dependências

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para verificar se uma porta está aberta
check_port() {
    local host=$1
    local port=$2
    local service=$3
    
    if nc -z -w5 "$host" "$port" 2>/dev/null || timeout 5 bash -c "cat < /dev/null > /dev/tcp/$host/$port" 2>/dev/null; then
        echo -e "${GREEN}✓${NC} $service is ${GREEN}UP${NC} on port $port"
        return 0
    else
        echo -e "${RED}✗${NC} $service is ${RED}DOWN${NC} on port $port"
        return 1
    fi
}

# Função para verificar HTTP endpoint
check_http() {
    local url=$1
    local service=$2
    
    status_code=$(curl -s -o /dev/null -w "%{http_code}" "$url" 2>/dev/null || echo "000")
    
    if [ "$status_code" -eq 200 ] || [ "$status_code" -eq 204 ]; then
        echo -e "${GREEN}✓${NC} $service HTTP endpoint is ${GREEN}healthy${NC} (HTTP $status_code)"
        return 0
    else
        echo -e "${RED}✗${NC} $service HTTP endpoint is ${RED}unhealthy${NC} (HTTP $status_code)"
        return 1
    fi
}

# Função para verificar gRPC endpoint
check_grpc() {
    local host=$1
    local port=$2
    local service=$3
    
    if grpcurl -plaintext -max-time 5 "$host:$port" list &>/dev/null; then
        echo -e "${GREEN}✓${NC} $service gRPC is ${GREEN}healthy${NC}"
        return 0
    else
        echo -e "${YELLOW}⚠${NC} $service gRPC check ${YELLOW}skipped${NC} (grpcurl not installed)"
        return 2
    fi
}

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Microservices Health Check${NC}"
echo -e "${BLUE}========================================${NC}\n"

total_checks=0
passed_checks=0
failed_checks=0
skipped_checks=0

# Infraestrutura
echo -e "${YELLOW}Infrastructure Services:${NC}"

check_port "localhost" "3306" "MySQL"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_port "localhost" "5672" "RabbitMQ"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_http "http://localhost:15672" "RabbitMQ Management"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

echo ""

# Monitoring
echo -e "${YELLOW}Monitoring Services:${NC}"

check_http "http://localhost:9090/-/healthy" "Prometheus"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_http "http://localhost:3000/api/health" "Grafana"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_port "localhost" "16686" "Jaeger UI"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

echo ""

# Microserviços
echo -e "${YELLOW}Microservices:${NC}"

# Order Service
check_port "localhost" "50051" "Order Service (gRPC)"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_http "http://localhost:50051/metrics" "Order Service (Metrics)"
((total_checks++))
result=$?
[ $result -eq 0 ] && ((passed_checks++))
[ $result -eq 1 ] && ((failed_checks++))
[ $result -eq 2 ] && ((skipped_checks++))

# Payment Service
check_port "localhost" "50052" "Payment Service (gRPC)"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_http "http://localhost:50052/metrics" "Payment Service (Metrics)"
((total_checks++))
result=$?
[ $result -eq 0 ] && ((passed_checks++))
[ $result -eq 1 ] && ((failed_checks++))
[ $result -eq 2 ] && ((skipped_checks++))

# User Service
check_port "localhost" "50053" "User Service (gRPC)"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_http "http://localhost:50053/metrics" "User Service (Metrics)"
((total_checks++))
result=$?
[ $result -eq 0 ] && ((passed_checks++))
[ $result -eq 1 ] && ((failed_checks++))
[ $result -eq 2 ] && ((skipped_checks++))

# Notification Service
check_port "localhost" "50054" "Notification Service (gRPC)"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

# Catalog Service
check_port "localhost" "50055" "Catalog Service (gRPC)"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

# BFF GraphQL
check_port "localhost" "8080" "BFF GraphQL"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

check_http "http://localhost:8080/graphql" "BFF GraphQL Endpoint"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

echo ""

# Frontend
echo -e "${YELLOW}Frontend:${NC}"

check_port "localhost" "3001" "React Frontend"
((total_checks++))
[ $? -eq 0 ] && ((passed_checks++)) || ((failed_checks++))

echo ""

# Resumo
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Health Check Summary${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "Total Checks:   $total_checks"
echo -e "${GREEN}Passed:${NC}         $passed_checks"
echo -e "${RED}Failed:${NC}         $failed_checks"
echo -e "${YELLOW}Skipped:${NC}        $skipped_checks"
echo ""

# Calcula percentual de sucesso
if [ $total_checks -gt 0 ]; then
    success_rate=$(( (passed_checks * 100) / total_checks ))
    
    if [ $success_rate -eq 100 ]; then
        echo -e "${GREEN}🎉 All services are healthy! ($success_rate%)${NC}"
        exit 0
    elif [ $success_rate -ge 80 ]; then
        echo -e "${YELLOW}⚠ Most services are healthy ($success_rate%)${NC}"
        exit 0
    else
        echo -e "${RED}❌ Multiple services are down ($success_rate%)${NC}"
        exit 1
    fi
else
    echo -e "${RED}❌ No services checked${NC}"
    exit 1
fi
