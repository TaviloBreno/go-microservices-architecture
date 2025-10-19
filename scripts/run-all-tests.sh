#!/bin/bash

# 🧪 Script para rodar todos os testes localmente antes de fazer push
# Este script simula o que o CI fará no GitHub Actions

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}╔══════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║  🧪 Running All Tests for Microservices Architecture   ║${NC}"
echo -e "${BLUE}╚══════════════════════════════════════════════════════════╝${NC}"
echo ""

# Array de serviços Go
GO_SERVICES=("order" "payment" "user" "notification" "catalog" "bff-graphql")

# Contador de testes
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0

# Função para rodar testes de um serviço
run_service_tests() {
    local service=$1
    echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${BLUE}📦 Testing service: ${GREEN}$service${NC}"
    echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    
    cd "services/$service"
    
    # Instalar dependências
    echo -e "${BLUE}📥 Installing dependencies...${NC}"
    go mod download
    go mod tidy
    
    # Rodar testes
    echo -e "${BLUE}🧪 Running unit tests...${NC}"
    if go test ./... -v -race -coverprofile=coverage.out -covermode=atomic; then
        echo -e "${GREEN}✅ Tests passed for $service${NC}"
        PASSED_TESTS=$((PASSED_TESTS + 1))
        
        # Gerar relatório de cobertura
        echo -e "${BLUE}📊 Coverage report:${NC}"
        go tool cover -func=coverage.out | tail -n 1
    else
        echo -e "${RED}❌ Tests failed for $service${NC}"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    cd ../..
    echo ""
}

# Rodar testes para todos os serviços Go
for service in "${GO_SERVICES[@]}"; do
    if [ -d "services/$service" ]; then
        run_service_tests "$service"
    else
        echo -e "${YELLOW}⚠️  Service directory not found: $service${NC}"
    fi
done

# Testar Frontend
echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}📦 Testing: ${GREEN}Frontend React${NC}"
echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

if [ -d "frontend" ]; then
    cd frontend
    
    echo -e "${BLUE}📥 Installing dependencies...${NC}"
    npm ci
    
    echo -e "${BLUE}🔍 Running linter...${NC}"
    npm run lint || echo -e "${YELLOW}⚠️  Linting issues found${NC}"
    
    echo -e "${BLUE}🏗️  Building frontend...${NC}"
    if npm run build; then
        echo -e "${GREEN}✅ Frontend build successful${NC}"
        PASSED_TESTS=$((PASSED_TESTS + 1))
    else
        echo -e "${RED}❌ Frontend build failed${NC}"
        FAILED_TESTS=$((FAILED_TESTS + 1))
    fi
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    cd ..
fi

# Validar Docker Compose
echo ""
echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
echo -e "${BLUE}🐳 Validating Docker Compose${NC}"
echo -e "${YELLOW}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"

if docker-compose config > /dev/null 2>&1; then
    echo -e "${GREEN}✅ Docker Compose configuration is valid${NC}"
    PASSED_TESTS=$((PASSED_TESTS + 1))
else
    echo -e "${RED}❌ Docker Compose configuration is invalid${NC}"
    FAILED_TESTS=$((FAILED_TESTS + 1))
fi

TOTAL_TESTS=$((TOTAL_TESTS + 1))

# Sumário final
echo ""
echo -e "${BLUE}╔══════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║                    📊 Test Summary                       ║${NC}"
echo -e "${BLUE}╚══════════════════════════════════════════════════════════╝${NC}"
echo -e "Total Tests:  ${BLUE}$TOTAL_TESTS${NC}"
echo -e "Passed:       ${GREEN}$PASSED_TESTS${NC}"
echo -e "Failed:       ${RED}$FAILED_TESTS${NC}"
echo ""

if [ $FAILED_TESTS -eq 0 ]; then
    echo -e "${GREEN}🎉 All tests passed! Ready to push.${NC}"
    exit 0
else
    echo -e "${RED}💥 Some tests failed. Please fix before pushing.${NC}"
    exit 1
fi