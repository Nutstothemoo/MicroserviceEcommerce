package orders

import (
	"errors"
	"microservice/pkg/common/price"
)

type ProductID string

var ErrEmptyProductID = errors.New("product id can not be empty")

type Product struct {
	id ProductID
	name string
	price price.Price	
}

func (p *Product) ProductID() ProductID {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p Product) Price() price.Price {
	return p.price
}

func NewProduct(id ProductID, name string, price price.Price) (Product, error) {	
	if id == "" {
		return Product{}, ErrEmptyProductID
	}
	if name == "" {
		return Product{}, errors.New("name can not be empty")
	}

	
	return Product{id, name, price}, nil
}

