package main 

import (
	"log"
	"fmt"
	"os"
	"net/http"
)

func main() {	
	 log.Println("Starting the payement microservice...")
	 defer log.Println("Shutting down the payement microservice...")
	 ctx := cmd.Context()

	 payementsInterface :=createPayementsMicroService()
	
	 if err := payementsInterface.Start(ctx); err != nil {
		 panic(err)
	}

}


func createPayementsMicroService() amqp.payementsInterface {
	cmd.WaitForService(os.Getenv("SHOP_RABBITMQ_ADDR"))
	payementsService := payements_app.NewPayementsService(
		payements_infra_orders.NewHTTPClient(os.Getenv("SHOP_ORDERS_SERVICE_ADDR")),
	)
	payementsInterface, err := amqp.NewPayementsInterface(
		fmt.Sprintf("amqp://%s", os.Getenv("SHOP_RABBITMQ_ADDR")),
		os.Getenv("SHOP_PAYEMENTS_SERVICE_QUEUE"),
		payementsService,
	)
	if err != nil {
		panic(err)
	}
	return payementsInterface
}