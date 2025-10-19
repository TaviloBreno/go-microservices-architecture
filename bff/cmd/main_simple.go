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

	// Handler para /notifications (versão simplificada para teste)
	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📢 Requisição de notificações: %s %s", r.Method, r.URL.Path)
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		
		// Resposta mock para testar
		mockResponse := []map[string]interface{}{
			{
				"id":         1,
				"payment_id": 1,
				"order_id":   123,
				"message":    "Payment processed successfully",
				"status":     "SENT",
				"created_at": "2025-10-19 17:12:30",
			},
		}
		
		json.NewEncoder(w).Encode(mockResponse)
	})

	// Handler para raiz (deve ser último)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🏠 Requisição raiz: %s %s", r.Method, r.URL.Path)
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("🚀 BFF Online - Microservices Architecture!"))
	})

	log.Println("🎯 Iniciando servidor HTTP na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}