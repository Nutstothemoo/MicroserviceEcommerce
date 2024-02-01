package main 

import (
	"log"
	"fmt"
	"os"
	"net/http"
	"cmd"
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

func createShopMicroservice () *chi.Mux {
	
	shopProductRepo := shop_infra_product.NewMemoryRepository(
	r := cmd.CreateRouter()
	
	shop_interface_public_http.AddRoutes(r, shopProductRepo)
	shop_interface_private_http.AddRoutes(r, shopProductRepo)	
	)
}