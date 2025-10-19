package messaging

import (
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/seu-usuario/go-microservices-architecture/services/order/internal/config"
)

// RabbitMQConnection gerencia a conexão com o RabbitMQ
type RabbitMQConnection struct {
	connection *amqp091.Connection
	channel    *amqp091.Channel
	queueName  string
}

// NewRabbitMQConnection cria uma nova conexão com RabbitMQ
func NewRabbitMQConnection() (*RabbitMQConnection, error) {
	// Ler configurações do ambiente
	user := config.GetEnv("RABBITMQ_USER", "guest")
	password := config.GetEnv("RABBITMQ_PASSWORD", "guest")
	host := config.GetEnv("RABBITMQ_HOST", "rabbitmq")
	port := config.GetEnv("RABBITMQ_PORT", "5672")
	queueName := config.GetEnv("RABBITMQ_QUEUE", "orders")

	// Construir URL de conexão
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, port)

	log.Printf("🐰 Conectando ao RabbitMQ: %s@%s:%s", user, host, port)

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

	log.Println("✅ Conectado ao RabbitMQ com sucesso")

	// Criar canal
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("falha ao abrir canal RabbitMQ: %w", err)
	}

	log.Println("📡 Canal RabbitMQ criado com sucesso")

	// Declarar fila
	_, err = ch.QueueDeclare(
		queueName, // nome da fila
		true,      // durable - persiste quando o servidor reinicia
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("falha ao declarar fila '%s': %w", queueName, err)
	}

	log.Printf("🎯 Fila '%s' declarada com sucesso", queueName)

	return &RabbitMQConnection{
		connection: conn,
		channel:    ch,
		queueName:  queueName,
	}, nil
}

// GetChannel retorna o canal RabbitMQ
func (r *RabbitMQConnection) GetChannel() *amqp091.Channel {
	return r.channel
}

// GetQueueName retorna o nome da fila
func (r *RabbitMQConnection) GetQueueName() string {
	return r.queueName
}

// Close fecha a conexão com RabbitMQ
func (r *RabbitMQConnection) Close() error {
	if r.channel != nil {
		if err := r.channel.Close(); err != nil {
			log.Printf("⚠️ Erro ao fechar canal RabbitMQ: %v", err)
		}
	}
	if r.connection != nil {
		if err := r.connection.Close(); err != nil {
			log.Printf("⚠️ Erro ao fechar conexão RabbitMQ: %v", err)
			return err
		}
	}
	log.Println("🔒 Conexão RabbitMQ fechada com sucesso")
	return nil
}
