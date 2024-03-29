package intraprocess

import (
	"microservice/pkg/orders/application"
	"microservice/pkg/orders/domain"
)
type OrdersInterface struct {
	service    application.OrdersService
}

func NewOrdersInterface(service application.OrdersService) *OrdersInterface {	
	return &OrdersInterface{service}
}

func (p OrdersInterface) MarkOrderAsPaid(orderID orders.OrderID ){
		cmd := application.MarkOrderAsPaidCommand{
			OrderID: orderID,
	}
	p.service.MarkOrderAsPaid(cmd)
}
