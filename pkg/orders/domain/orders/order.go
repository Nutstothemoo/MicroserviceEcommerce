package orders

import (
	"errors"
)

type ID string 

var ErrEmptyOrderID = errors.New("order id can not be empty")

type Order struct {
	id ID
	product Product
	adress Adress
	paid bool
}

func (o *Order) ID() ID {
	return o.id
}

func (o *Order) Product() Product {
	return o.product
}

func (o *Order) MarkAsPaid() {
	o.paid = true
}

func (o *Order) Adress() Adress {
	return o.adress
}

func (o *Order) Paid() bool {
	return o.paid
}

func NewOrder(id ID, product Product, adress Adress) (*Order, error) {
	if id == "" {
		return nil, ErrEmptyOrderID
	}
	if product == nil {
		return nil, errors.New("product can not be empty")
	}
	if adress == nil {
		return nil, errors.New("adress can not be empty")
	}
	return &Order{id, product, adress, false}, nil
}
