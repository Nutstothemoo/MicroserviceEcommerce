package application

import (
	"log"
	"github.com/Nutstothemoo/microservice/pkg/common/"
	"github.com/pkg/errors"
)

type productsService interface {

}

type payementsService interface {
}

type 	OrdersService struct {
}

func NewOrdersService(productsService productsService, payementsService payementsService) *OrdersService {
	return &OrdersService{}
}
type PlaceOrderCommand struct {
}


type MarkOrderAsPaidCommand struct {
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {
return nil
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	return nil
}

func()OrderById(id string) (Order, error) {
	return Order{}, nil
}