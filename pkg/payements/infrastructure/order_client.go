package payements_infra_orders

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	orders "microservice/pkg/orders/domain"
)

type OrdersClient struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewOrdersClient() (*OrdersClient, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
			return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
			return nil, err
	}

	q, err := ch.QueueDeclare(
			"orders", // name
			false,    // durable
			false,    // delete when unused
			false,    // exclusive
			false,    // no-wait
			nil,      // arguments
	)
	if err != nil {
			return nil, err
	}

	return &OrdersClient{
			Conn:    conn,
			Channel: ch,
			Queue:   q,
	}, nil
}

func (c *OrdersClient) GetOrder() (*orders.Order, error) {
	msgs, err := c.Channel.Consume(
			c.Queue.Name, // queue
			"",           // consumer
			true,         // auto-ack
			false,        // exclusive
			false,        // no-local
			false,        // no-wait
			nil,          // args
	)
	if err != nil {
			return nil, err
	}

	msg := <-msgs

	var order orders.Order
	err = json.Unmarshal(msg.Body, &order)
	if err != nil {
			return nil, err
	}

	return &order, nil
}