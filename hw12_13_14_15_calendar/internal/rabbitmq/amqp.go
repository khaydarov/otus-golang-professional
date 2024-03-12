package rabbitmq

import (
	"errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(rmqUrl string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(rmqUrl)
	if err != nil {
		return nil, errors.New("failed to connect to rabbitmq")
	}

	return conn, nil
}
