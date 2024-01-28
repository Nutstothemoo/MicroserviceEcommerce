package products

import (
	"errors"
)

var ErrNotFound = errors.New("Product not found")

type Repository interface {
	Save(*product) error
	ByID(ID) (*product, error)
}