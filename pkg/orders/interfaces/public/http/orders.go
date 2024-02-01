package http

import (
	"net/http"
	common_http "microservice/pkg/common/http"
	"microservice/pkg/orders/application"
	"microservice/pkg/orders/domain"
	"github.com/google/uuid"
	// "microservice/pkg/products/domain/products"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func AddRoutes(router *chi.Mux, service application.OrdersService, repository orders.Repository){
	resource := OrdersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders/{id}/paid", resource.GetPaid)
}

type OrdersResource struct {
	service    application.OrdersService
	repository orders.Repository
}

type PostOrderRequest struct {
	ProductID string `json:"product_id"`
	Address PostOrderAddress `json:"address"`
}

type PostOrderAddress struct {
	Name string       `json:"name"`
	City string 			`json:"city"`
	Postcode string 	`json:"postcode"`
	Street string 		`json:"street"`
	Country string 		`json:"country"`
}

func (o OrdersResource) Post(w http.ResponseWriter, r *http.Request) {
	req := PostOrderRequest{}
	if err := render.Decode(r , &req); err != nil {
		_ = render.Render(w, r, common_http.ErrBadRequest(err))
		return
	}
	cmd := application.PostOrderCommand{
		OrderID: orders.OrderID(uuid.New().String()),
		ProductID: products_domain.ProductID(req.ProductID),
		Address: application.PlaceOrderCommandAdress(req.Address),
	}
	if err:= o.service.PlaceOrder(cmd); err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	w.WriteHeader(http.StatusCreated)
	render.JSON(w, r, PostOrderResponse{
		OrderID: string(cmd.OrderID)})
}

type PostOrderResponse struct {
	OrderID string 	`json:"order_id"`
}

type OrderPaidView struct {
	ID string 		`json:"id"`
	isPaid bool 	`json:"is_paid"`
}

func (o OrdersResource) GetPaid(w http.ResponseWriter, r *http.Request) {
	orderID := orders.OrderID(chi.URLParam(r, "id"))
	order, err := o.repository.Get(orderID)
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	render.Respond(w, r, OrderPaidView{
		ID: string(order.ID()),
		isPaid: order.	IsPaid(),
	})
}
