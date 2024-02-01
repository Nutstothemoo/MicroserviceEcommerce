package application

import (
    payements_infra_orders "microservice/pkg/payements/infrastructure/orders"
)

// PayementsService represents the application service for payments.
type PayementsService struct {
    ordersClient payements_infra_orders.OrdersClient
}

func NewPayementsService(ordersClient payements_infra_orders.OrdersClient) *PayementsService {
    return &PayementsService{
        ordersClient: ordersClient,
    }
}

// func (s *PayementsService) ProcessPayment(orderID string) error {

// }