package orders

import (
	"context"
)

type Repository interface {
	Create(order *Order) error
	GetById(id string) (Order, error)
	Update(ctx context.Context, order *Order) error
	Delete(ctx context.Context, id string) error
	MarkAsPaid(ctx context.Context, id string) error
}