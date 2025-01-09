package commons

import (
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

var rabbitMQInstance *RabbitMQ
var once sync.Once

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func InitRabbitMQ(url string) (*RabbitMQ, error) {
	var err error
	once.Do(func() {
		rabbitMQInstance, err = createRabbitMQInstance(url)
	})
	return rabbitMQInstance, err
}

func createRabbitMQInstance(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func GetRabbitMQInstance() *RabbitMQ {
	return rabbitMQInstance
}

func (r *RabbitMQ) Publish(exchange, routingKey, body string) error {
	err := r.Channel.Publish(
		exchange,   // Exchange
		routingKey, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return err
}

func (r *RabbitMQ) Close() {
	if err := r.Channel.Close(); err != nil {
		log.Printf("Failed to close channel: %s", err)
	}
	if err := r.Conn.Close(); err != nil {
		log.Printf("Failed to close connection: %s", err)
	}
}

func (r *RabbitMQ) ExchangeDeclare(exchangeName, exchangeType string, durable, autoDelete, internal, noWait bool) error {
	err := r.Channel.ExchangeDeclare(
		exchangeName,
		exchangeType,
		durable,
		autoDelete,
		internal,
		noWait,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare exchange: %w", err)
	}
	return nil
}

func (r *RabbitMQ) QueueDeclare(queueName string, durable, autoDelete, exclusive, noWait bool) (amqp.Queue, error) {
	q, err := r.Channel.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		exclusive,
		noWait,
		nil,
	)
	if err != nil {
		return q, fmt.Errorf("failed to declare queue: %w", err)
	}
	return q, nil
}
