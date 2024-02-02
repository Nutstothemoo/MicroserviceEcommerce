package main

import (
	"log"
	"microservice/pkg/common/cmd"
	application_shop "microservice/pkg/shop/application"
	application_order "microservice/pkg/orders/application"
	order_private_http "microservice/pkg/orders/interfaces/private/http"
	orders_public_http "microservice/pkg/orders/interfaces/public/http"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting the order microservice...")

	ctx := cmd.Context()
	r, closeFn := createOrderMicroservice()
	defer closeFn()

	server := &http.Server{
		Addr:    os.Getenv("SHOP_ORDER_SERVICE_ADDR"),
		Handler: r,
	}
	go func() {
		if err:= server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
		
	}()
	<-ctx.Done()
	log.Println("Shutting down the order microservice...")
	if err := server.Close(); err != nil {
		panic(err)
	}
}


func createOrderMicroservice() (router *chi.Mux, closeFn func()){
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	shopHTTPClient := orders_infra_product.NewHTTPClient(os.Getenv("SHOP_SERVICE_ADDR"))

	r:= cmd.CreateRouter()

	// Initialize your products service, payments service, and orders repository here.
	productsService := 	application_shop.NewProductsService(productsRepository, productReadModel)
	payementsService := application_shop.NewPayementsService(shopHTTPClient)
	// ordersRepository := application_order.NewOrdersRepository()
	
	// Initialize your orders service.

	ordersService := application_order.NewOrdersService(productsService, payementsService, ordersRepository)

	orders_public_http.AddRoutes(r, *ordersService, ordersRepo )
	order_private_http.AddRoutes(r, *ordersService, ordersRepo)

	return r, func() {
				shopHTTPClient.Close()
	}
}
