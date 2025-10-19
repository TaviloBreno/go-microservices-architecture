package publisher

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"payment-service/internal/config"
	"payment-service/internal/domain"

	"github.com/rabbitmq/amqp091-go"
)

// PaymentPublisher gerencia a publicação de eventos de pagamento
type PaymentPublisher struct {
	conn      *amqp091.Connection
	channel   *amqp091.Channel
	queueName string
}

// PaymentEvent representa o evento de pagamento a ser enviado
type PaymentEvent struct {
	PaymentID uint    `json:"payment_id"`
	OrderID   uint    `json:"order_id"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
}

// NewPaymentPublisher cria uma nova instância do publisher
func NewPaymentPublisher() (*PaymentPublisher, error) {
	cfg := config.LoadConfig()

	// Construir URL de conexão
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort)

	// Conectar com retry
	var conn *amqp091.Connection
	var err error

	for i := 0; i < 5; i++ {
		conn, err = amqp091.Dial(url)
		if err == nil {
			break
		}
		log.Printf("⚠️  Tentativa %d: Erro ao conectar RabbitMQ Publisher: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("falha ao conectar com RabbitMQ após 5 tentativas: %w", err)
	}

	log.Println("✅ Conectado ao RabbitMQ Publisher com sucesso")

	// Abrir canal
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("falha ao abrir canal RabbitMQ: %w", err)
	}

	log.Println("📡 Canal RabbitMQ Publisher criado com sucesso")

	// Declarar a fila payments
	paymentQueueName := "payments"
	_, err = ch.QueueDeclare(
		paymentQueueName, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("falha ao declarar fila %s: %w", paymentQueueName, err)
	}

	log.Printf("🎯 Fila '%s' declarada com sucesso", paymentQueueName)

	return &PaymentPublisher{
		conn:      conn,
		channel:   ch,
		queueName: paymentQueueName,
	}, nil
}

// PublishPaymentEvent publica um evento de pagamento na fila
func (p *PaymentPublisher) PublishPaymentEvent(payment *domain.Payment) error {
	// Criar evento
	event := PaymentEvent{
		PaymentID: payment.ID,
		OrderID:   payment.OrderID,
		Status:    payment.Status,
		Amount:    payment.Amount,
		CreatedAt: payment.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	// Serializar para JSON
	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("falha ao serializar evento de pagamento: %w", err)
	}

	// Publicar mensagem
	err = p.channel.Publish(
		"",          // exchange
		p.queueName, // routing key (queue name)
		false,       // mandatory
		false,       // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return fmt.Errorf("falha ao publicar evento de pagamento: %w", err)
	}

	log.Printf("📤 Evento de pagamento publicado na fila '%s':", p.queueName)
	log.Printf("   🆔 Payment ID: %d", event.PaymentID)
	log.Printf("   📝 Order ID: %d", event.OrderID)
	log.Printf("   📊 Status: %s", event.Status)
	log.Printf("   💰 Valor: R$ %.2f", event.Amount)
	log.Printf("   ⏰ Criado em: %s", event.CreatedAt)

	return nil
}

// Close fecha as conexões do publisher
func (p *PaymentPublisher) Close() error {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.conn != nil {
		p.conn.Close()
	}
	log.Println("🔌 Conexões RabbitMQ Publisher fechadas")
	return nil
}
