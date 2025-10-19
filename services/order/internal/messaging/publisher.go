package messaging

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/domain"
)

// OrderPublisher gerencia a publicação de eventos de pedidos
type OrderPublisher struct {
	rabbitmqConn *RabbitMQConnection
}

// OrderEvent representa o evento de criação de pedido
type OrderEvent struct {
	ID        uint    `json:"id"`
	Customer  string  `json:"customer"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	EventType string  `json:"event_type"`
}

// NewOrderPublisher cria uma nova instância do publisher de pedidos
func NewOrderPublisher(rabbitmqConn *RabbitMQConnection) *OrderPublisher {
	return &OrderPublisher{
		rabbitmqConn: rabbitmqConn,
	}
}

// PublishOrderCreated publica um evento de pedido criado na fila RabbitMQ
func (p *OrderPublisher) PublishOrderCreated(order *domain.Order) error {
	// Criar contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Converter pedido para evento
	orderEvent := OrderEvent{
		ID:        order.ID,
		Customer:  order.Customer,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Price:     order.Price,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		EventType: "order.created",
	}

	// Serializar para JSON
	messageBody, err := json.Marshal(orderEvent)
	if err != nil {
		log.Printf("❌ Erro ao serializar pedido para JSON: %v", err)
		return err
	}

	log.Printf("📤 Preparando para enviar mensagem: %s", string(messageBody))

	// Publicar mensagem
	err = p.rabbitmqConn.GetChannel().PublishWithContext(
		ctx,
		"",                            // exchange
		p.rabbitmqConn.GetQueueName(), // routing key (nome da fila)
		false,                         // mandatory
		false,                         // immediate
		amqp091.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp091.Persistent, // torna a mensagem persistente
			Timestamp:    time.Now(),
			Body:         messageBody,
		},
	)

	if err != nil {
		log.Printf("❌ Erro ao publicar mensagem no RabbitMQ: %v", err)
		return err
	}

	log.Printf("📨 Mensagem enviada para fila RabbitMQ com ID: %d", order.ID)
	return nil
}

// PublishOrderUpdated publica um evento de pedido atualizado (para uso futuro)
func (p *OrderPublisher) PublishOrderUpdated(order *domain.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	orderEvent := OrderEvent{
		ID:        order.ID,
		Customer:  order.Customer,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Price:     order.Price,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
		EventType: "order.updated",
	}

	messageBody, err := json.Marshal(orderEvent)
	if err != nil {
		log.Printf("❌ Erro ao serializar pedido atualizado para JSON: %v", err)
		return err
	}

	err = p.rabbitmqConn.GetChannel().PublishWithContext(
		ctx,
		"",
		p.rabbitmqConn.GetQueueName(),
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp091.Persistent,
			Timestamp:    time.Now(),
			Body:         messageBody,
		},
	)

	if err != nil {
		log.Printf("❌ Erro ao publicar mensagem de atualização no RabbitMQ: %v", err)
		return err
	}

	log.Printf("📨 Mensagem de atualização enviada para fila RabbitMQ com ID: %d", order.ID)
	return nil
}
