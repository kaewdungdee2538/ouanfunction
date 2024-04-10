package amqp

import (
	"github.com/rabbitmq/amqp091-go"
)

type Ampq interface {
	Connection(url string) (*amqp091.Connection, error)
	Channel(conn *amqp091.Connection) (*amqp091.Channel, error)
	SetupQueueTypeDirectExchange(ch *amqp091.Channel,  queueName string,exchangeName string, routingKey string) error 
	SetupQueueTypeTopic(ch *amqp091.Channel, queueName string,  exchangeName string,routingHeadKey string) error 
}
