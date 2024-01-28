package http

import (
	"net/http"
)

func AddRoutes(){
	resource := OrdersResource{service, repository}
	router.Post("/orders", resource.Post)
	router.Get("/orders", resource.GetPaid)
}

type OrdersResource struct {
	service    application.OrdersService
	repository orders.Repository
}