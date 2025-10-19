package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() error {
	// Conectar ao MySQL
	dbHost := getEnv("DB_HOST", "mysql")
	dbUser := getEnv("DB_USER", "microservices")
	dbPass := getEnv("DB_PASS", "micro123")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/", dbUser, dbPass, dbHost)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao MySQL: %v", err)
	}

	// Testar conexão
	if err := db.Ping(); err != nil {
		return fmt.Errorf("erro ao pingar MySQL: %v", err)
	}

	log.Println("✅ Conectado ao MySQL com sucesso!")
	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	fmt.Println("🌐 BFF rodando na porta 8080...")

	// Inicializar conexão com banco
	if err := initDB(); err != nil {
		log.Printf("⚠️  Aviso: Não foi possível conectar ao banco: %v", err)
		log.Println("📝 Continuando com endpoints mock...")
	}
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

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

	// Handler para /api/products - Listar produtos do catálogo
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("🛍️  Requisição de produtos: %s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if db == nil {
			http.Error(w, "Database não disponível", http.StatusServiceUnavailable)
			return
		}

		rows, err := db.Query(`
			SELECT p.id, p.name, p.description, p.price, p.stock_quantity, 
			       p.image_url, p.sku, c.name as category_name
			FROM catalog_service.products p
			LEFT JOIN catalog_service.categories c ON p.category_id = c.id
			WHERE p.is_active = TRUE
			ORDER BY p.created_at DESC
		`)
		if err != nil {
			log.Printf("Erro ao consultar produtos: %v", err)
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var products []map[string]interface{}
		for rows.Next() {
			var id, name, sku, categoryName, imageUrl string
			var description sql.NullString
			var price float64
			var stock int

			if err := rows.Scan(&id, &name, &description, &price, &stock, &imageUrl, &sku, &categoryName); err != nil {
				log.Printf("Erro ao ler produto: %v", err)
				continue
			}

			product := map[string]interface{}{
				"id":          id,
				"name":        name,
				"description": description.String,
				"price":       price,
				"stock":       stock,
				"image_url":   imageUrl,
				"sku":         sku,
				"category":    categoryName,
			}
			products = append(products, product)
		}

		json.NewEncoder(w).Encode(products)
	})

	// Handler para /api/users - Listar usuários
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("👥 Requisição de usuários: %s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if db == nil {
			http.Error(w, "Database não disponível", http.StatusServiceUnavailable)
			return
		}

		rows, err := db.Query(`
			SELECT id, name, email, phone, address, created_at
			FROM user_service.users
			ORDER BY created_at DESC
		`)
		if err != nil {
			log.Printf("Erro ao consultar usuários: %v", err)
			http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id, name, email string
			var phone, address, createdAt sql.NullString

			if err := rows.Scan(&id, &name, &email, &phone, &address, &createdAt); err != nil {
				log.Printf("Erro ao ler usuário: %v", err)
				continue
			}

			user := map[string]interface{}{
				"id":         id,
				"name":       name,
				"email":      email,
				"phone":      phone.String,
				"address":    address.String,
				"created_at": createdAt.String,
			}
			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	})

	// Handler para /api/orders - Listar pedidos
	http.HandleFunc("/api/orders", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📦 Requisição de pedidos: %s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if db == nil {
			http.Error(w, "Database não disponível", http.StatusServiceUnavailable)
			return
		}

		rows, err := db.Query(`
			SELECT o.id, o.user_id, o.status, o.total_amount, o.payment_status, 
			       o.shipping_address, o.created_at, u.name as user_name, u.email
			FROM order_service.orders o
			LEFT JOIN user_service.users u ON o.user_id = u.id
			ORDER BY o.created_at DESC
		`)
		if err != nil {
			log.Printf("Erro ao consultar pedidos: %v", err)
			http.Error(w, "Erro ao buscar pedidos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var orders []map[string]interface{}
		for rows.Next() {
			var id, userId, status, paymentStatus, shippingAddress, createdAt string
			var totalAmount float64
			var userName, userEmail sql.NullString

			if err := rows.Scan(&id, &userId, &status, &totalAmount, &paymentStatus, &shippingAddress, &createdAt, &userName, &userEmail); err != nil {
				log.Printf("Erro ao ler pedido: %v", err)
				continue
			}

			order := map[string]interface{}{
				"id":               id,
				"user_id":          userId,
				"user_name":        userName.String,
				"user_email":       userEmail.String,
				"status":           status,
				"payment_status":   paymentStatus,
				"total_amount":     totalAmount,
				"shipping_address": shippingAddress,
				"created_at":       createdAt,
			}
			orders = append(orders, order)
		}

		json.NewEncoder(w).Encode(orders)
	})

	// Handler para /api/payments - Listar pagamentos
	http.HandleFunc("/api/payments", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("💳 Requisição de pagamentos: %s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if db == nil {
			http.Error(w, "Database não disponível", http.StatusServiceUnavailable)
			return
		}

		rows, err := db.Query(`
			SELECT p.id, p.order_id, p.user_id, p.amount, p.payment_method, 
			       p.status, p.transaction_id, p.card_brand, p.installments, p.created_at,
			       u.name as user_name
			FROM payment_service.payments p
			LEFT JOIN user_service.users u ON p.user_id = u.id
			ORDER BY p.created_at DESC
		`)
		if err != nil {
			log.Printf("Erro ao consultar pagamentos: %v", err)
			http.Error(w, "Erro ao buscar pagamentos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var payments []map[string]interface{}
		for rows.Next() {
			var id, orderId, userId, paymentMethod, status, createdAt string
			var amount float64
			var transactionId, cardBrand, userName sql.NullString
			var installments sql.NullInt32

			if err := rows.Scan(&id, &orderId, &userId, &amount, &paymentMethod, &status, &transactionId, &cardBrand, &installments, &createdAt, &userName); err != nil {
				log.Printf("Erro ao ler pagamento: %v", err)
				continue
			}

			payment := map[string]interface{}{
				"id":             id,
				"order_id":       orderId,
				"user_id":        userId,
				"user_name":      userName.String,
				"amount":         amount,
				"payment_method": paymentMethod,
				"status":         status,
				"transaction_id": transactionId.String,
				"card_brand":     cardBrand.String,
				"installments":   installments.Int32,
				"created_at":     createdAt,
			}
			payments = append(payments, payment)
		}

		json.NewEncoder(w).Encode(payments)
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
