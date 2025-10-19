#!/bin/bash

# ========================================
# Docker Swarm Deployment Script
# ========================================

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
STACK_NAME="${STACK_NAME:-go-ms}"
VERSION="${VERSION:-latest}"
DOCKERHUB_USERNAME="${DOCKERHUB_USERNAME:-tavilobreno}"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Docker Swarm Deployment${NC}"
echo -e "${BLUE}========================================${NC}\n"

# Check if running on Swarm manager
if ! docker node ls &>/dev/null; then
    echo -e "${RED}❌ Error: This node is not a Swarm manager${NC}"
    echo -e "${YELLOW}Initialize Swarm with: docker swarm init${NC}"
    exit 1
fi

echo -e "${GREEN}✓${NC} Running on Swarm manager\n"

# Create secrets if they don't exist
echo -e "${YELLOW}Creating secrets...${NC}"

create_secret() {
    local secret_name=$1
    local secret_value=$2
    
    if ! docker secret inspect "$secret_name" &>/dev/null; then
        echo "$secret_value" | docker secret create "$secret_name" -
        echo -e "${GREEN}✓${NC} Created secret: $secret_name"
    else
        echo -e "${BLUE}ℹ${NC} Secret already exists: $secret_name"
    fi
}

# Prompt for secrets if not set in environment
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
    read -sp "Enter SMTP password (optional, press Enter to skip): " SMTP_PASSWORD
    echo
fi

# Create secrets
create_secret "mysql_root_password" "$MYSQL_ROOT_PASSWORD"
create_secret "mysql_password" "$MYSQL_PASSWORD"
create_secret "rabbitmq_password" "$RABBITMQ_PASSWORD"
create_secret "grafana_password" "$GRAFANA_PASSWORD"
[ -n "$SMTP_PASSWORD" ] && create_secret "smtp_password" "$SMTP_PASSWORD"

echo ""

# Validate stack configuration
echo -e "${YELLOW}Validating stack configuration...${NC}"
if docker-compose -f deployment/docker-swarm/stack.yml config > /dev/null 2>&1; then
    echo -e "${GREEN}✓${NC} Stack configuration is valid\n"
else
    echo -e "${RED}❌ Invalid stack configuration${NC}"
    exit 1
fi

# Deploy stack
echo -e "${YELLOW}Deploying stack: $STACK_NAME${NC}"
echo -e "${BLUE}Version: $VERSION${NC}"
echo -e "${BLUE}Registry: $DOCKERHUB_USERNAME${NC}\n"

export DOCKERHUB_USERNAME
export VERSION

docker stack deploy -c deployment/docker-swarm/stack.yml --with-registry-auth "$STACK_NAME"

echo -e "\n${GREEN}✓${NC} Stack deployed successfully\n"

# Wait for services to start
echo -e "${YELLOW}Waiting for services to start...${NC}"
sleep 10

# Show service status
echo -e "\n${BLUE}Service Status:${NC}"
docker stack services "$STACK_NAME"

# Check for failed services
echo -e "\n${YELLOW}Checking for failed services...${NC}"
FAILED_SERVICES=$(docker stack services "$STACK_NAME" --format '{{.Name}}: {{.Replicas}}' | grep '0/' || true)

if [ -n "$FAILED_SERVICES" ]; then
    echo -e "${RED}⚠ Warning: Some services have no replicas running:${NC}"
    echo "$FAILED_SERVICES"
    echo -e "\n${YELLOW}Check logs with: docker service logs <service-name>${NC}"
else
    echo -e "${GREEN}✓${NC} All services are running\n"
fi

# Show endpoints
echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Service Endpoints${NC}"
echo -e "${BLUE}========================================${NC}"
echo -e "${GREEN}Frontend:${NC}        http://localhost:3001"
echo -e "${GREEN}GraphQL API:${NC}     http://localhost:8080/graphql"
echo -e "${GREEN}Grafana:${NC}         http://localhost:3000 (admin/${GRAFANA_PASSWORD:-admin123})"
echo -e "${GREEN}Prometheus:${NC}      http://localhost:9090"
echo -e "${GREEN}Jaeger:${NC}          http://localhost:16686"
echo -e "${GREEN}RabbitMQ:${NC}        http://localhost:15672 (guest/${RABBITMQ_PASSWORD:-guest})"
echo ""

echo -e "${GREEN}🎉 Deployment completed!${NC}"
echo -e "${YELLOW}Run './deployment/scripts/health-check-swarm.sh' to verify all services${NC}"
