package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"payment-service/internal/config"
	"payment-service/internal/service"

	"github.com/rabbitmq/amqp091-go"
)

// OrderEvent representa o evento de pedido recebido do RabbitMQ
type OrderEvent struct {
	ID        uint    `json:"id"`
	Customer  string  `json:"customer"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	EventType string  `json:"event_type"`
}

// OrderConsumer gerencia o consumo de mensagens da fila de pedidos
type OrderConsumer struct {
	connection     *amqp091.Connection
	channel        *amqp091.Channel
	paymentService service.PaymentService
	queueName      string
}

// NewOrderConsumer cria uma nova instância do consumer
func NewOrderConsumer(paymentService service.PaymentService) (*OrderConsumer, error) {
	// Ler configurações do ambiente
	cfg := config.LoadConfig()

	// Construir URL de conexão
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort)

	log.Printf("🐰 Conectando ao RabbitMQ Consumer: %s@%s:%s", cfg.RabbitMQUser, cfg.RabbitMQHost, cfg.RabbitMQPort)

	// Conectar com retry logic
	var conn *amqp091.Connection
	var err error
	maxRetries := 10

	for i := 0; i < maxRetries; i++ {
		conn, err = amqp091.Dial(rabbitURL)
		if err == nil {
			break
		}
		log.Printf("⏳ Tentativa %d/%d de conexão RabbitMQ falhou: %v. Aguardando 3 segundos...", i+1, maxRetries, err)
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao RabbitMQ após %d tentativas: %w", maxRetries, err)
	}

	log.Println("✅ Conectado ao RabbitMQ Consumer com sucesso")

	// Criar canal
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("falha ao abrir canal RabbitMQ: %w", err)
	}

	log.Println("📡 Canal RabbitMQ Consumer criado com sucesso")

	// Configurar QoS para processar uma mensagem por vez
	err = ch.Qos(1, 0, false)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("falha ao configurar QoS: %w", err)
	}

	return &OrderConsumer{
		connection:     conn,
		channel:        ch,
		paymentService: paymentService,
		queueName:      cfg.RabbitMQQueue,
	}, nil
}

// ConsumeOrders inicia o consumo de mensagens da fila de pedidos
func (c *OrderConsumer) ConsumeOrders(ctx context.Context) error {
	log.Printf("🎧 Iniciando consumo da fila '%s'...", c.queueName)

	// Registrar como consumidor
	msgs, err := c.channel.Consume(
		c.queueName,                // queue
		"payment-service-consumer", // consumer tag
		false,                      // auto-ack (false para ack manual)
		false,                      // exclusive
		false,                      // no-local
		false,                      // no-wait
		nil,                        // args
	)
	if err != nil {
		return fmt.Errorf("falha ao registrar consumidor: %w", err)
	}

	log.Printf("✅ Consumer registrado com sucesso na fila '%s'", c.queueName)

	// Processar mensagens
	for {
		select {
		case <-ctx.Done():
			log.Println("🛑 Parando consumer por contexto cancelado")
			return ctx.Err()

		case delivery, ok := <-msgs:
			if !ok {
				log.Println("🔚 Canal de mensagens fechado")
				return nil
			}

			c.processMessage(delivery)
		}
	}
}

// processMessage processa uma mensagem individual
func (c *OrderConsumer) processMessage(delivery amqp091.Delivery) {
	log.Printf("📨 Nova mensagem recebida da fila %s", c.queueName)
	log.Printf("📝 Conteúdo da mensagem: %s", string(delivery.Body))

	// Parse do JSON
	var orderEvent OrderEvent
	if err := json.Unmarshal(delivery.Body, &orderEvent); err != nil {
		log.Printf("❌ Erro ao fazer parse da mensagem JSON: %v", err)
		// Rejeitar mensagem com requeue=false (vai para DLQ se configurada)
		delivery.Nack(false, false)
		return
	}

	// Validar tipo de evento
	if orderEvent.EventType != "order.created" {
		log.Printf("⚠️ Tipo de evento ignorado: %s", orderEvent.EventType)
		delivery.Ack(false)
		return
	}

	// Calcular total (preço * quantidade)
	totalAmount := orderEvent.Price * float64(orderEvent.Quantity)

	log.Printf("💳 Processando pagamento - OrderID: %d, Amount: %.2f",
		orderEvent.ID, totalAmount)

	// Processar pagamento
	payment, err := c.paymentService.ProcessPayment(orderEvent.ID, totalAmount)
	if err != nil {
		log.Printf("❌ Erro ao processar pagamento para OrderID %d: %v", orderEvent.ID, err)
		// Rejeitar com requeue=true para tentar novamente
		delivery.Nack(false, true)
		return
	}

	log.Printf("✅ Pagamento salvo com sucesso no banco - ID: %d, Status: %s",
		payment.ID, payment.Status)

	// Confirmar processamento da mensagem
	if err := delivery.Ack(false); err != nil {
		log.Printf("⚠️ Erro ao confirmar mensagem: %v", err)
	}
}

// Close fecha as conexões do consumer
func (c *OrderConsumer) Close() error {
	if c.channel != nil {
		if err := c.channel.Close(); err != nil {
			log.Printf("⚠️ Erro ao fechar canal RabbitMQ Consumer: %v", err)
		}
	}
	if c.connection != nil {
		if err := c.connection.Close(); err != nil {
			log.Printf("⚠️ Erro ao fechar conexão RabbitMQ Consumer: %v", err)
			return err
		}
	}
	log.Println("🔒 Conexão RabbitMQ Consumer fechada com sucesso")
	return nil
}
