package http

import (	
	"net/http"
	common_http "microservice/pkg/common/http"
	"microservice/pkg/orders/application"
	"microservice/pkg/orders/domain/orders"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func AddRoutes(r *chi.Mux, service application.OrdersService, ordersRepository orders.Repository) {
	resource := OrdersResource{service, ordersRepository}
	r.Post("/orders/{id}/paid", resource.PostPaid)
} 


type OrdersResource struct {
	service application.OrdersService
	ordersRepository orders.Repository
}

func (o OrdersResource) PostPaid(w http.ResponseWriter, r *http.Request) {
	orderID := chi.URLParam(r, "id")
	cmd := application.MarkOrderAsPaidCommand{
		OrderID: orders.ID(orderID),
	}
	if err := o.service.MarkOrderAsPaid(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	render.NoContent(w, r)
}
