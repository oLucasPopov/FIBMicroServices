package service

import (
	"context"
	"delivery/config"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type RabbitMQSender struct {
	config config.RabbitMQ
}

func MakeRabbitMQSender() *RabbitMQSender {
	c, err := config.MakeRabbitMQ()
	if err != nil {
		log.Panic(err.Error())
	}
	return &RabbitMQSender{config: *c}
}

func (rmq *RabbitMQSender) CloseConnection() {
	if err := rmq.config.Conn.Close(); err != nil {
		fmt.Println(err)
	}
}

func (rmq *RabbitMQSender) ConvertAndSend(data any) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ch, err := rmq.config.Conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	q, err := ch.QueueDeclare(
		"order",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonBytes,
		},
	)

	if err != nil {
		return err
	}

	return nil

}
