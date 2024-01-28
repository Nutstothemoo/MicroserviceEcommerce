package application

import (
	"log"
	"microservice/pkg/common/price"
	"microservice/pkg/orders/domain/orders"
	"github.com/pkg/errors"
)

type productsService interface {
	ProductsByID(id orders.ProductID) (orders.Product, error)
}


type payementsService interface {
	InitializeOrderPayement(orderID orders.ID, price price.Price) (error)
}

type 	OrdersService struct {
	productsService productsService
	payementsService payementsService
	ordersRepository orders.Repository
}

func NewOrdersService(productsService productsService, payementsService payementsService, ordersRepository orders.Repository) *OrdersService {
	return &OrdersService{
		productsService: productsService,
		payementsService: payementsService,
		ordersRepository: ordersRepository,
	}
}
type PlaceOrderCommand struct {
		OrderID string
		ProductID string
}
type PlaceOrderCommandAddress struct {
	Name string
	Street string
	City string
	PostalCode string
	Country string
}

type MarkOrderAsPaidCommand struct {
	OrderID orders.ID	
}

func (s OrdersService) PlaceOrder(cmd PlaceOrderCommand) error {
	address, err := orders.NewAddress(
		cmd.Address.Name,
		cmd.Address.Street, 
		cmd.Address.City, 
		cmd.Address.PostalCode,
		cmd.Address.Country,
	)
	if err != nil {
		return errors.Wrap(err, "Error creating address")
	}
	product, err := s.productsService.ProductsByID(cmd.ProductID)
	if err != nil {
		return errors.Wrap(err, "Error getting product by id")
	}
	newOrder, err := orders.NewOrder(
		orders.ID(cmd.OrderID),
		product,
		address,
	)
	if err != nil {
		return errors.Wrap(err, "Error creating new order")
	}
	if err := s.ordersRepository.Save(newOrder); err != nil {
		return errors.Wrap(err, "Error saving new order")
	}
	if err := s.payementsService.InitializeOrderPayement(newOrder.ID(), newOrder.Product().Price()); err != nil {
		return errors.Wrap(err, "Error initializing order payement")
	}
	log.Printf("Order %s has been placed", NewOrder.ID)
	return nil
}	

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	o , err := s.OrderById(cmd.OrderID)
	if err != nil {
		return errors.Wrap(err, "Error getting order by id")
	}
	o.MarkAsPaid()
	if err := s.ordersRepository.Save(o); err != nil {
		return errors.Wrap(err, "Error saving order")
	}
	log.Printf("Order %s has been marked as paid", o.ID())
	return nil
}

func(s OrdersService ) OrderById(id orders.ID) (orders.Order, error) {
	o, err := s.ordersRepository.OrderById(id)
	if err != nil {
		return orders.Order{}, errors.Wrap(err, "Error getting order by id")
	}
	return o, nil
}