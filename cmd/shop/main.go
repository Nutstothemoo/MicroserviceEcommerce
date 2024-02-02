package main 

import (
	"log"
	"os"
	"net/http"
	"microservice/pkg/common/cmd"
	shop_infra_product "microservice/pkg/shop/infrastructure/products"
	shop_interface_public_http "microservice/pkg/shop/interfaces/public/http"
	shop_interface_private_http "microservice/pkg/shop/interfaces/private/http"
	"github.com/go-chi/chi"
)

func main() {
	log.Println("Starting the shop microservice...")
	ctx:= cmd.Context()
	r, closeFn := createShopMicroservice()

	server:= &http.Server{
		Addr: os.Getenv("SHOP_SERVICE_BIND_ADDR"),
		Handler: r,
	}	
	go func() {
		if err:= server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
		<-ctx.Done()
		log.Println("Shutting down the shop microservice...")
		if err:= server.Close(); err != nil {
			panic(err)
		}

}

func createShopMicroservice () (*chi.Mux, func())  {
	
	shopProductRepo := shop_infra_product.NewMemoryRepository()
	r := cmd.CreateRouter()
	
	shop_interface_public_http.AddRoutes(r, shopProductRepo)
	shop_interface_private_http.AddRoutes(r, shopProductRepo)	
	closeFn := func() {}

	return r, closeFn
}