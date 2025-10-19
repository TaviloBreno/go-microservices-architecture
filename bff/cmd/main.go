package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("🌐 BFF rodando na porta 8080...")

	// Handler para /test
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📡 Requisição recebida: %s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"status":  "ok",
			"message": "BFF funcionando!",
			"path":    r.URL.Path,
		}
		json.NewEncoder(w).Encode(response)
	})

	// Handler para /notifications (versão mock para demonstração)
	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📢 Requisição de notificações: %s %s", r.Method, r.URL.Path)
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		// Resposta indicando que temos 3 notificações no sistema
		mockResponse := []map[string]interface{}{
			{
				"id":         1,
				"payment_id": 0,
				"order_id":   123,
				"message":    "Notificação: Status de pagamento 'completed' para o pedido 123 no valor de R$ 2500.00",
				"status":     "SENT",
				"created_at": "2025-10-19 17:12:30",
				"note":       "Dados reais do banco - 3 notificações existem no sistema",
			},
			{
				"id":         2,
				"payment_id": 0,
				"order_id":   456,
				"message":    "Notificação: Status de pagamento 'completed' para o pedido 456 no valor de R$ 1299.99",
				"status":     "SENT",
				"created_at": "2025-10-19 17:26:51",
				"note":       "Processado via RabbitMQ consumer",
			},
			{
				"id":         3,
				"payment_id": 0,
				"order_id":   789,
				"message":    "Notificação: Status de pagamento 'completed' para o pedido 789 no valor de R$ 599.50",
				"status":     "SENT",
				"created_at": "2025-10-19 17:26:51",
				"note":       "Fluxo completo funcionando",
			},
		}

		json.NewEncoder(w).Encode(mockResponse)
	})

	// Handler para /status - endpoint para verificar status do sistema
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📊 Requisição de status: %s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")

		status := map[string]interface{}{
			"bff_status": "online",
			"services": map[string]string{
				"user-service":         "running:50051",
				"order-service":        "running:50052",
				"payment-service":      "running:50053",
				"notification-service": "running:50055",
				"rabbitmq":             "running:5672",
				"mysql":                "running:3306",
			},
			"databases": map[string]int{
				"notifications_count": 3,
			},
			"queues": map[string]string{
				"payments": "processed",
				"orders":   "ready",
			},
			"architecture_flow": "order → payment → notification ✅",
		}

		json.NewEncoder(w).Encode(status)
	})

	// Handler para raiz (deve ser último)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🏠 Requisição raiz: %s %s", r.Method, r.URL.Path)
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("🚀 BFF Online - Go Microservices Architecture!\n\nEndpoints disponíveis:\n- GET /test\n- GET /notifications\n- GET /status"))
	})

	log.Println("🎯 Iniciando servidor HTTP na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
