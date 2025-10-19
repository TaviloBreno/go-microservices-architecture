#!/bin/bash

# ============================================================================
# ENVIRONMENT CHECK SCRIPT
# ============================================================================
# Verifica se o ambiente está pronto para executar o projeto
# ============================================================================

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}============================================${NC}"
echo -e "${BLUE}  Verificando Ambiente...${NC}"
echo -e "${BLUE}============================================${NC}"
echo ""

# Variável para track de erros
ERRORS=0
WARNINGS=0

# Função para verificar comando
check_command() {
    if command -v $1 &> /dev/null; then
        VERSION=$($1 --version 2>&1 | head -n 1)
        echo -e "${GREEN}✓${NC} $1 instalado: $VERSION"
        return 0
    else
        echo -e "${RED}✗${NC} $1 não encontrado"
        ERRORS=$((ERRORS + 1))
        return 1
    fi
}

# Função para verificar porta
check_port() {
    PORT=$1
    SERVICE=$2
    
    if command -v lsof &> /dev/null; then
        if lsof -Pi :$PORT -sTCP:LISTEN -t >/dev/null 2>&1; then
            echo -e "${YELLOW}⚠${NC} Porta $PORT ($SERVICE) já em uso"
            WARNINGS=$((WARNINGS + 1))
            return 1
        else
            echo -e "${GREEN}✓${NC} Porta $PORT ($SERVICE) disponível"
            return 0
        fi
    elif command -v netstat &> /dev/null; then
        if netstat -tuln | grep ":$PORT " >/dev/null 2>&1; then
            echo -e "${YELLOW}⚠${NC} Porta $PORT ($SERVICE) já em uso"
            WARNINGS=$((WARNINGS + 1))
            return 1
        else
            echo -e "${GREEN}✓${NC} Porta $PORT ($SERVICE) disponível"
            return 0
        fi
    else
        echo -e "${YELLOW}⚠${NC} Não foi possível verificar porta $PORT (lsof/netstat não disponível)"
        WARNINGS=$((WARNINGS + 1))
        return 1
    fi
}

# Função para verificar espaço em disco
check_disk_space() {
    AVAILABLE=$(df -h . | tail -1 | awk '{print $4}' | sed 's/G//')
    
    if [ -z "$AVAILABLE" ]; then
        echo -e "${YELLOW}⚠${NC} Não foi possível verificar espaço em disco"
        WARNINGS=$((WARNINGS + 1))
        return 1
    fi
    
    # Converte para número (remove possível 'G')
    AVAILABLE_NUM=$(echo $AVAILABLE | sed 's/[^0-9.]//g')
    
    if (( $(echo "$AVAILABLE_NUM < 10" | bc -l) )); then
        echo -e "${YELLOW}⚠${NC} Espaço em disco baixo: ${AVAILABLE}G disponível (recomendado: 10GB+)"
        WARNINGS=$((WARNINGS + 1))
        return 1
    else
        echo -e "${GREEN}✓${NC} Espaço em disco: ${AVAILABLE}G disponível"
        return 0
    fi
}

# Verificar comandos obrigatórios
echo -e "\n${BLUE}Verificando dependências obrigatórias:${NC}"
check_command docker
check_command docker-compose || check_command docker compose

# Verificar comandos opcionais
echo -e "\n${BLUE}Verificando dependências opcionais (desenvolvimento):${NC}"
check_command go || echo -e "${YELLOW}  (opcional - apenas para desenvolvimento Go)${NC}"
check_command node || echo -e "${YELLOW}  (opcional - apenas para desenvolvimento Frontend)${NC}"
check_command make || echo -e "${YELLOW}  (opcional - facilita uso de comandos)${NC}"

# Verificar Docker daemon
echo -e "\n${BLUE}Verificando Docker:${NC}"
if docker info >/dev/null 2>&1; then
    echo -e "${GREEN}✓${NC} Docker daemon rodando"
    
    # Verificar memória do Docker
    DOCKER_MEM=$(docker info 2>/dev/null | grep "Total Memory" | awk '{print $3}')
    if [ ! -z "$DOCKER_MEM" ]; then
        echo -e "${GREEN}✓${NC} Memória Docker: ${DOCKER_MEM}"
        
        # Converter para GB e verificar
        MEM_GB=$(echo $DOCKER_MEM | sed 's/GiB//')
        if (( $(echo "$MEM_GB < 4" | bc -l) )); then
            echo -e "${YELLOW}⚠${NC} Memória Docker baixa (recomendado: 8GB+)"
            WARNINGS=$((WARNINGS + 1))
        fi
    fi
else
    echo -e "${RED}✗${NC} Docker daemon não está rodando"
    echo -e "${YELLOW}  Execute: sudo systemctl start docker (Linux) ou inicie Docker Desktop (Windows/Mac)${NC}"
    ERRORS=$((ERRORS + 1))
fi

# Verificar portas
echo -e "\n${BLUE}Verificando portas:${NC}"
check_port 3306 "MySQL"
check_port 5672 "RabbitMQ"
check_port 15672 "RabbitMQ Management"
check_port 8080 "BFF GraphQL"
check_port 3001 "Frontend"
check_port 9090 "Prometheus"
check_port 3000 "Grafana"
check_port 16686 "Jaeger"

# Verificar espaço em disco
echo -e "\n${BLUE}Verificando recursos:${NC}"
check_disk_space

# Verificar arquivo .env
echo -e "\n${BLUE}Verificando configuração:${NC}"
if [ -f ".env" ]; then
    echo -e "${GREEN}✓${NC} Arquivo .env encontrado"
else
    echo -e "${YELLOW}⚠${NC} Arquivo .env não encontrado"
    echo -e "${YELLOW}  Execute: cp .env.example .env${NC}"
    WARNINGS=$((WARNINGS + 1))
fi

# Verificar docker-compose.yml
if [ -f "docker-compose.yml" ]; then
    echo -e "${GREEN}✓${NC} docker-compose.yml encontrado"
else
    echo -e "${RED}✗${NC} docker-compose.yml não encontrado"
    ERRORS=$((ERRORS + 1))
fi

# Resumo final
echo -e "\n${BLUE}============================================${NC}"
echo -e "${BLUE}  Resumo${NC}"
echo -e "${BLUE}============================================${NC}"

if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
    echo -e "${GREEN}✓ Ambiente pronto para uso!${NC}"
    echo -e "\nPróximos passos:"
    echo -e "  1. docker-compose up -d"
    echo -e "  2. Aguarde 30-60 segundos"
    echo -e "  3. Acesse: http://localhost:3001"
    exit 0
elif [ $ERRORS -eq 0 ]; then
    echo -e "${YELLOW}⚠ Ambiente pronto com $WARNINGS aviso(s)${NC}"
    echo -e "\nVocê pode prosseguir, mas considere resolver os avisos."
    echo -e "\nPróximos passos:"
    echo -e "  1. docker-compose up -d"
    echo -e "  2. Aguarde 30-60 segundos"
    echo -e "  3. Acesse: http://localhost:3001"
    exit 0
else
    echo -e "${RED}✗ Encontrados $ERRORS erro(s) e $WARNINGS aviso(s)${NC}"
    echo -e "\nCorreja os erros antes de prosseguir."
    echo -e "Consulte a documentação: docs/INSTALLATION.md"
    exit 1
fi
