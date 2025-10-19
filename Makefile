.PHONY: help test lint build run clean docker-build docker-up docker-down coverage

# Cores para output
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## Mostra esta mensagem de ajuda
help:
	@echo ''
	@echo 'Uso:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "  ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)

##@ Testes

test: ## Roda todos os testes
	@echo "${GREEN}Running all tests...${RESET}"
	@./scripts/run-all-tests.sh

test-order: ## Testa apenas Order Service
	@echo "${GREEN}Testing Order Service...${RESET}"
	@cd services/order && go test ./... -v

test-payment: ## Testa apenas Payment Service
	@echo "${GREEN}Testing Payment Service...${RESET}"
	@cd services/payment && go test ./... -v

test-user: ## Testa apenas User Service
	@echo "${GREEN}Testing User Service...${RESET}"
	@cd services/user && go test ./... -v

test-frontend: ## Testa Frontend React
	@echo "${GREEN}Testing Frontend...${RESET}"
	@cd frontend && npm test

coverage: ## Gera relatório de cobertura
	@echo "${GREEN}Generating coverage reports...${RESET}"
	@for service in order payment user notification catalog; do \
		echo "${CYAN}Coverage for $$service:${RESET}"; \
		cd services/$$service && go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out | tail -n 1; \
		cd ../..; \
	done

##@ Linting

lint: ## Roda linter em todos os serviços
	@echo "${GREEN}Running linter...${RESET}"
	@for service in order payment user notification catalog bff-graphql; do \
		echo "${CYAN}Linting $$service:${RESET}"; \
		cd services/$$service && golangci-lint run || true; \
		cd ../..; \
	done

lint-fix: ## Roda linter com auto-fix
	@echo "${GREEN}Running linter with auto-fix...${RESET}"
	@for service in order payment user notification catalog bff-graphql; do \
		echo "${CYAN}Linting $$service:${RESET}"; \
		cd services/$$service && golangci-lint run --fix || true; \
		cd ../..; \
	done

##@ Build

build: ## Builda todos os serviços
	@echo "${GREEN}Building all services...${RESET}"
	@for service in order payment user notification catalog bff-graphql; do \
		echo "${CYAN}Building $$service:${RESET}"; \
		cd services/$$service/cmd && go build -o ../$$service-service; \
		cd ../../..; \
	done

build-frontend: ## Builda o frontend
	@echo "${GREEN}Building frontend...${RESET}"
	@cd frontend && npm run build

##@ Docker

docker-build: ## Builda todas as imagens Docker
	@echo "${GREEN}Building Docker images...${RESET}"
	@docker-compose build

docker-up: ## Sobe todos os containers
	@echo "${GREEN}Starting containers...${RESET}"
	@docker-compose up -d

docker-down: ## Para todos os containers
	@echo "${GREEN}Stopping containers...${RESET}"
	@docker-compose down

docker-logs: ## Mostra logs dos containers
	@docker-compose logs -f

docker-ps: ## Lista containers rodando
	@docker-compose ps

docker-restart: docker-down docker-up ## Reinicia todos os containers

health-check: ## Verifica saúde de todos os serviços
	@echo "${GREEN}Checking services health...${RESET}"
	@bash scripts/health-check.sh

##@ Desenvolvimento

dev-order: ## Roda Order Service em modo dev
	@echo "${GREEN}Starting Order Service...${RESET}"
	@cd services/order/cmd && go run main.go

dev-payment: ## Roda Payment Service em modo dev
	@echo "${GREEN}Starting Payment Service...${RESET}"
	@cd services/payment/cmd && go run main.go

dev-user: ## Roda User Service em modo dev
	@echo "${GREEN}Starting User Service...${RESET}"
	@cd services/user/cmd && go run main.go

dev-bff: ## Roda BFF GraphQL em modo dev
	@echo "${GREEN}Starting BFF GraphQL...${RESET}"
	@cd services/bff-graphql/cmd && go run main.go

dev-frontend: ## Roda Frontend em modo dev
	@echo "${GREEN}Starting Frontend...${RESET}"
	@cd frontend && npm run dev

##@ Database

db-migrate: ## Roda migrations do banco
	@echo "${GREEN}Running database migrations...${RESET}"
	@docker-compose up -d mysql
	@sleep 5
	@echo "${CYAN}Migrations completed${RESET}"

db-reset: ## Reseta o banco de dados
	@echo "${YELLOW}Resetting database...${RESET}"
	@docker-compose down -v
	@docker-compose up -d mysql
	@sleep 10
	@echo "${GREEN}Database reset completed${RESET}"

##@ Monitoring

monitoring-up: ## Sobe apenas serviços de monitoramento
	@echo "${GREEN}Starting monitoring services...${RESET}"
	@docker-compose up -d prometheus grafana jaeger

monitoring-logs: ## Mostra logs do monitoramento
	@docker-compose logs -f prometheus grafana jaeger

prometheus: ## Abre Prometheus no browser
	@echo "${GREEN}Opening Prometheus...${RESET}"
	@open http://localhost:9090 || xdg-open http://localhost:9090 || start http://localhost:9090

grafana: ## Abre Grafana no browser
	@echo "${GREEN}Opening Grafana (admin/admin123)...${RESET}"
	@open http://localhost:3000 || xdg-open http://localhost:3000 || start http://localhost:3000

jaeger: ## Abre Jaeger no browser
	@echo "${GREEN}Opening Jaeger...${RESET}"
	@open http://localhost:16686 || xdg-open http://localhost:16686 || start http://localhost:16686

##@ Cleanup

clean: ## Limpa arquivos de build e cache
	@echo "${GREEN}Cleaning up...${RESET}"
	@find . -name "*.exe" -type f -delete
	@find . -name "coverage.out" -type f -delete
	@find . -name "*.log" -type f -delete
	@go clean -cache -testcache -modcache
	@echo "${CYAN}Cleanup completed${RESET}"

clean-docker: ## Remove todos os containers e volumes
	@echo "${YELLOW}Removing all containers and volumes...${RESET}"
	@docker-compose down -v --remove-orphans
	@docker system prune -f
	@echo "${GREEN}Docker cleanup completed${RESET}"

##@ CI/CD

ci-local: ## Simula CI pipeline localmente
	@echo "${GREEN}Running CI pipeline locally...${RESET}"
	@./scripts/run-all-tests.sh

release: ## Cria uma nova release (usage: make release VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then \
		echo "${YELLOW}Please specify VERSION (e.g., make release VERSION=v1.0.0)${RESET}"; \
		exit 1; \
	fi
	@echo "${GREEN}Creating release $(VERSION)...${RESET}"
	@git tag $(VERSION)
	@git push origin $(VERSION)
	@echo "${CYAN}Release $(VERSION) created and pushed!${RESET}"

##@ Production Deployment

deploy-swarm: ## Deploy para Docker Swarm
	@echo "${GREEN}Deploying to Docker Swarm...${RESET}"
	@bash deployment/scripts/deploy-swarm.sh

deploy-k8s: ## Deploy para Kubernetes
	@echo "${GREEN}Deploying to Kubernetes...${RESET}"
	@bash deployment/scripts/deploy-k8s.sh

rollback-swarm: ## Rollback no Docker Swarm
	@echo "${YELLOW}Rolling back Docker Swarm deployment...${RESET}"
	@bash deployment/scripts/rollback-swarm.sh

rollback-k8s: ## Rollback no Kubernetes
	@echo "${YELLOW}Rolling back Kubernetes deployment...${RESET}"
	@bash deployment/scripts/rollback-k8s.sh

health-check-swarm: ## Health check do Docker Swarm
	@echo "${GREEN}Checking Docker Swarm health...${RESET}"
	@bash deployment/scripts/health-check-swarm.sh

health-check-k8s: ## Health check do Kubernetes
	@echo "${GREEN}Checking Kubernetes health...${RESET}"
	@bash deployment/scripts/health-check-k8s.sh

##@ Default

.DEFAULT_GOAL := help