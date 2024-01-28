package orders

import (
	"errors"
	"pkg/common/price"
)

type productID string
var ErrEmptyProductID = errors.New("product id can not be empty")

type Product struct {
	id productID
	name string
	price price.Price	
}

func (p *Product) ID() productID {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() price.Price {
	return p.price
}

func NewProduct(id productID, name string, price price.Price) (Product, error) {	
	if id == "" {
		return nil, ErrEmptyProductID
	}
	if name == "" {
		return nil, errors.New("name can not be empty")
	}

	
	return Product{id, name, price}, nil
}
