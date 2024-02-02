package application

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	orders "microservice/pkg/orders/domain"
	payements_infra_orders "microservice/pkg/payements/infrastructure"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PayementsService struct {
    OrdersClient *payements_infra_orders.OrdersClient
}

type PaymentDetails struct {
    OrderID string  `json:"order_id"`
    Amount  float64 `json:"amount"`
}

func NewPayementsService(ordersClient *payements_infra_orders.OrdersClient) *PayementsService {
    return &PayementsService{
        OrdersClient: ordersClient,
    }
}

func (s *PayementsService) ProcessPayment(orderID string) error {
    order, err := s.OrdersClient.GetOrder()
    if err != nil {
        return err
    }

    err = makePayment(order)
    if err != nil {
        return err
    }

    err = sendPaymentMessage(orderID)
    if err != nil {
        return err
    }

    return nil
}

func sendPaymentMessage(orderID string) error {

    // Connect to RabbitMQ server
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        return err
    }
    defer conn.Close()

    // Create a channel
    ch, err := conn.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "payments", // name
        false,      // durable
        false,      // delete when unused
        false,      // exclusive
        false,      // no-wait
        nil,        // arguments
    )
    if err != nil {
        return err
    }

    body := "Payment for order " + orderID + " has been processed"
    ctx := context.Background()

    err = ch.PublishWithContext(
        ctx,    // context
        "",     // exchange
        q.Name, // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    return err
}

func makePayment(order *orders.Order) error {

    paymentDetails := PaymentDetails{
        OrderID: string(order.OrderID()),
        Amount:  1,
    }

    jsonData, err := json.Marshal(paymentDetails)
    if err != nil {
        return err
    }


    resp, err := http.Post("https://payment-api.example.com/pay", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("payment API returned status code %d", resp.StatusCode)
    }

    return nil
}