#!/bin/bash

# ========================================
# Docker Swarm Rollback Script
# ========================================

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

STACK_NAME="${STACK_NAME:-go-ms}"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Docker Swarm Rollback${NC}"
echo -e "${BLUE}========================================${NC}\n"

echo -e "${RED}⚠️  WARNING: This will rollback all services to their previous version${NC}"
read -p "Are you sure you want to continue? (yes/no): " confirm

if [ "$confirm" != "yes" ]; then
    echo -e "${YELLOW}Rollback cancelled${NC}"
    exit 0
fi

echo -e "\n${YELLOW}Rolling back services...${NC}\n"

# Get all services in the stack
SERVICES=$(docker stack services $STACK_NAME --format '{{.Name}}')

for SERVICE in $SERVICES; do
    echo -e "${BLUE}Rolling back: $SERVICE${NC}"
    
    if docker service rollback $SERVICE; then
        echo -e "${GREEN}✓${NC} Rolled back: $SERVICE"
    else
        echo -e "${RED}❌ Failed to rollback: $SERVICE${NC}"
    fi
    
    sleep 2
done

echo -e "\n${YELLOW}Waiting for services to stabilize...${NC}"
sleep 20

# Show service status
echo -e "\n${BLUE}Service Status After Rollback:${NC}"
docker stack services $STACK_NAME

echo -e "\n${GREEN}Rollback completed!${NC}"
echo -e "${YELLOW}Check service logs if any issues persist${NC}"
