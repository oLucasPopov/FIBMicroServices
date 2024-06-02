package workers

import (
	"cook/config"
	"cook/service"
	"fmt"
	"log"
)

type RabbitMQListener struct {
	config config.RabbitMQ
}

func MakeRabbitMQListener() *RabbitMQListener {
	c, err := config.MakeRabbitMQ()
	if err != nil {
		log.Panic(err.Error())
	}
	return &RabbitMQListener{config: *c}
}

func (rmq *RabbitMQListener) Listen() <-chan error {
	forever := make(chan error, 1)

	ch, err := rmq.config.Conn.Channel()
	if err != nil {
		fmt.Println(err.Error())
		forever <- err
		return forever
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
		forever <- err
		return forever
	}

	if err := ch.Qos(1, 0, false); err != nil {
		fmt.Println(err.Error())
		forever <- err
		return forever
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		fmt.Println(err.Error())
		forever <- err
		return forever
	}

	go func() {
		for d := range msgs {
			orderService := service.NewOrderService{}
			err := orderService.Handle(d.Body)
			if err != nil {
				forever <- err
				fmt.Println(forever)
			} else {
				if err := d.Ack(false); err != nil {
					forever <- err
					fmt.Println(forever)
				} else {
					fmt.Println("message acknowledged", string(d.Body))
				}
			}
		}
	}()

	return forever

}
