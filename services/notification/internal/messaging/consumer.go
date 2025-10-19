package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"notification-service/internal/config"
	"notification-service/internal/service"

	"github.com/rabbitmq/amqp091-go"
)

// PaymentEvent representa o evento de pagamento recebido do RabbitMQ
type PaymentEvent struct {
	PaymentID uint    `json:"payment_id"`
	OrderID   uint    `json:"order_id"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
}

// PaymentConsumer gerencia o consumo de mensagens da fila payments
type PaymentConsumer struct {
	conn                *amqp091.Connection
	channel             *amqp091.Channel
	notificationService service.NotificationService
	queueName           string
}

// NewPaymentConsumer cria uma nova instância do consumer
func NewPaymentConsumer(notificationService service.NotificationService) (*PaymentConsumer, error) {
	// Ler configurações do ambiente
	cfg := config.LoadConfig()

	// Construir URL de conexão
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort)

	log.Printf("🐰 Conectando ao RabbitMQ Consumer: %s@%s:%s", cfg.RabbitMQUser, cfg.RabbitMQHost, cfg.RabbitMQPort)

	// Conectar com retry logic
	var conn *amqp091.Connection
	var err error

	for i := 0; i < 10; i++ {
		conn, err = amqp091.Dial(rabbitURL)
		if err == nil {
			break
		}
		log.Printf("⚠️  Tentativa %d: Erro ao conectar RabbitMQ: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar com RabbitMQ após 10 tentativas: %w", err)
	}

	log.Println("✅ Conectado ao RabbitMQ Consumer com sucesso")

	// Abrir canal
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

	// Declarar a fila
	_, err = ch.QueueDeclare(
		cfg.RabbitMQQueue, // name
		true,              // durable
		false,             // delete when unused
		false,             // exclusive
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("falha ao declarar fila %s: %w", cfg.RabbitMQQueue, err)
	}

	log.Printf("🎯 Fila '%s' declarada com sucesso", cfg.RabbitMQQueue)

	return &PaymentConsumer{
		conn:                conn,
		channel:             ch,
		notificationService: notificationService,
		queueName:           cfg.RabbitMQQueue,
	}, nil
}

// ConsumePayments inicia o consumo de mensagens da fila payments
func (c *PaymentConsumer) ConsumePayments(ctx context.Context) error {
	log.Printf("🎧 Iniciando consumo da fila '%s'...", c.queueName)

	// Registrar consumer
	msgs, err := c.channel.Consume(
		c.queueName, // queue
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return fmt.Errorf("falha ao registrar consumer: %w", err)
	}

	log.Printf("✅ Consumer registrado com sucesso na fila '%s'", c.queueName)

	// Processar mensagens
	go func() {
		for delivery := range msgs {
			select {
			case <-ctx.Done():
				log.Println("🛑 Contexto cancelado, parando consumer...")
				return
			default:
				c.processMessage(delivery)
			}
		}
	}()

	log.Printf("⚡ Consumer aguardando mensagens da fila '%s'...", c.queueName)
	return nil
}

// processMessage processa uma mensagem individual
func (c *PaymentConsumer) processMessage(delivery amqp091.Delivery) {
	log.Printf("📨 Nova mensagem recebida da fila %s", c.queueName)
	log.Printf("📄 Tamanho da mensagem: %d bytes", len(delivery.Body))

	var paymentEvent PaymentEvent
	err := json.Unmarshal(delivery.Body, &paymentEvent)
	if err != nil {
		log.Printf("❌ Erro ao fazer unmarshal da mensagem: %v", err)
		log.Printf("📋 Conteúdo da mensagem: %s", string(delivery.Body))

		// Rejeitar mensagem com erro de formato
		delivery.Nack(false, false)
		return
	}

	log.Printf("📦 Evento de pagamento deserializado:")
	log.Printf("   🆔 Payment ID: %d", paymentEvent.PaymentID)
	log.Printf("   📝 Order ID: %d", paymentEvent.OrderID)
	log.Printf("   📊 Status: %s", paymentEvent.Status)
	log.Printf("   💰 Valor: R$ %.2f", paymentEvent.Amount)
	log.Printf("   ⏰ Criado em: %s", paymentEvent.CreatedAt)

	// Processar notificação
	err = c.notificationService.ProcessPaymentNotification(
		paymentEvent.PaymentID,
		paymentEvent.OrderID,
		paymentEvent.Status,
		paymentEvent.Amount,
	)

	if err != nil {
		log.Printf("❌ Erro ao processar notificação de pagamento: %v", err)

		// Rejeitar mensagem e reenviar para retry
		delivery.Nack(false, true)
		return
	}

	// Acknowleda mensagem como processada com sucesso
	delivery.Ack(false)
	log.Printf("✅ Mensagem processada e confirmada com sucesso")
}

// Close fecha as conexões do consumer
func (c *PaymentConsumer) Close() error {
	if c.channel != nil {
		c.channel.Close()
	}
	if c.conn != nil {
		c.conn.Close()
	}
	log.Println("🔌 Conexões RabbitMQ fechadas")
	return nil
}
