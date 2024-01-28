package products

import (
	"errors"
	"price"
)

type ID string

var (
	ErrEmpty ID = errors.New("ID cannot be empty")
	ErrEmpty Name = errors.New("Name cannot be empty")
)

type Product struct {
	id          ID
	name        string
	description string
	price       price.Price
}

func NewProduct(id ID, name string, price price.Price, description string) (Product, error) {
	if id == "" {
		return nil, ErrEmpty
	}
	if name == "" {
		return nil, ErrEmpty
	}
	return &Product{
		id:          id,
		name:        name,
		price:       price,
		description: description,
	}, nil
}

func (p Product) ID() ID {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) Price() price.Price {
	return p.price
}

func (p Product) Description() string {
	return p.description
}



