package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

func MakeRabbitMQ() (*RabbitMQ, error) {
	conn, err := amqp.Dial("amqps://piwqkjrn:FD4iBth90ANlx9Ehxe5nIdLCBhzEWJkB@jackal.rmq.cloudamqp.com/piwqkjrn")

	if err != nil {
		return nil, err
	}

	rabbitMQ := &RabbitMQ{
		Conn: conn,
	}

	return rabbitMQ, nil
}
