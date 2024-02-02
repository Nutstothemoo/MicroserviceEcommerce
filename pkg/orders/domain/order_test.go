package orders

import (
    "testing"
		"microservice/pkg/common/price"
)

func TestNewOrder(t *testing.T) {
    id := OrderID("1")
    product := Product{} // Remplacez ceci par une instance valide de Product
    adress := Adress{}  // Remplacez ceci par une instance valide de Adress

    order, err := NewOrder(id, product, &adress)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if order.OrderID() != id {
        t.Errorf("Expected ID to be %v, got %v", id, order.OrderID())
    }

    if order.Product() != product {
        t.Errorf("Expected Product to be %v, got %v", product, order.Product())
    }

    if order.Adress() != adress {
        t.Errorf("Expected Adress to be %v, got %v", adress, order.Adress())
    }

    if order.IsPaid() != false {
        t.Errorf("Expected Paid to be false, got %v", order.IsPaid())
    }
}

func TestMarkAsPaid(t *testing.T) {
    id := OrderID("1")
    product := Product{
				id: "1",
				name: "name",
				price: *price.NewPriceP(100, "USD"),
		} 
    adress := Adress{}  
    order, err := NewOrder(id, product, &adress)
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    order.MarkAsPaid()

    if order.IsPaid() != true {
        t.Errorf("Expected Paid to be true, got %v", order.IsPaid())
    }
}