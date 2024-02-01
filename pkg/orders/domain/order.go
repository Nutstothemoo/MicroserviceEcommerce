package orders

import (
	"errors"
)

type OrderID string 

var ErrEmptyOrderID = errors.New("order id can not be empty")

type Order struct {
	id OrderID
	product Product
	adress Adress
	paid bool
}

func (o *Order) OrderID() OrderID {
	return o.id
}

func (o Order) Product() Product {
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

func NewOrder(id OrderID, product Product, adress *Adress) (*Order, error) {
	if id == "" {
		return nil, ErrEmptyOrderID
	}
	return &Order{id, product, *adress, false}, nil
}


