package queue

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

func InitRabbitMQ() *amqp.Connection {
	url := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		viper.GetString("RABBITMQ_USER"),
		viper.GetString("RABBITMQ_PASS"),
		viper.GetString("RABBITMQ_HOST"),
		viper.GetString("RABBITMQ_PORT"),
	)

	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ { // retry up to 10 times
		conn, err = amqp.Dial(url)
		if err == nil {
			fmt.Println("✅ Connected to RabbitMQ")
			return conn
		}
		fmt.Println("⏳ Waiting for RabbitMQ to be ready... (retrying)")
		time.Sleep(3 * time.Second)
	}

	panic("Failed to connect to RabbitMQ: " + err.Error())
}
