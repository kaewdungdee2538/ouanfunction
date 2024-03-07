package amqp

import (
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type amqpRabbit struct {
	ttl int
}

func NewAmqpRabbit(ttl int) amqpRabbit {
	return amqpRabbit{ttl}
}

func (amqp amqpRabbit) Connection(url string) (*amqp091.Connection, error) {
	// connect ampq
	var reconnectAttempts = 1
	for {
		conn, err := amqp091.Dial(url)
		if err == nil {
			log.Println("connection to rabbitMQ success")
			return conn, nil
		}

		log.Printf("Failed to connect to RabbitMQ: %v, round %d retrying...", err, reconnectAttempts)
		// Exponential backoff
		time.Sleep(time.Second * 2) // Double the wait time after each attempt
		reconnectAttempts++
	}
}

func (amqp amqpRabbit) Channel(conn *amqp091.Connection) (*amqp091.Channel, error) {
	// channel amqp
	ch, err := conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return nil, err
	}
	log.Println("channel rabbitMQ success")
	return ch, nil
}

func (amqp amqpRabbit) SetupQueueTypeDirectExchange(ch *amqp091.Channel, queueName string, exchangeName string, routingKey string) error {

	exchangeType := "direct"

	err := ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // exchange type
		true,         // durable
		false,        // auto delete
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare a exchange")
		return err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp091.Table{
			"x-message-ttl": amqp.ttl, // Set the TTL for the queue
		},
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue")
		return err
	}

	err = ch.QueueBind(
		q.Name,       // name
		routingKey,   // key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue binding")
		return err
	}

	log.Println("setup queue success")
	return nil
}

func (amqp amqpRabbit) SetupQueueTypeTopic(ch *amqp091.Channel, queueName string, exchangeName string, routingHeadKey string) error {

	exchangeType := "topic"
	keyBindingTopic := fmt.Sprintf("%s.%s", routingHeadKey, "*")

	err := ch.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // exchange type
		true,         // durable
		false,        // auto delete
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare a exchange")
		return err
	}

	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp091.Table{
			"x-message-ttl": amqp.ttl, // Set the TTL for the queue
		},
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue")
		return err
	}

	err = ch.QueueBind(
		q.Name,          // name
		keyBindingTopic, // key
		exchangeName,    // exchange
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare a queue binding")
		return err
	}

	log.Println("setup queue success")
	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
	}
}
