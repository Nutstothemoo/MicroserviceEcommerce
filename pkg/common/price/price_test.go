package price

import (
    "testing"
)

func TestNewPrice(t *testing.T) {
    _, err := NewPrice(0, "USD")
    if err != ErrPriceTooLow {
        t.Errorf("Expected ErrPriceTooLow, got %v", err)
    }

    _, err = NewPrice(100, "US")
    if err != ErrInvalidCurrency {
        t.Errorf("Expected ErrInvalidCurrency, got %v", err)
    }

    price, err := NewPrice(100, "USD")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if price.Cents() != 100 {
        t.Errorf("Expected 100 cents, got %v", price.Cents())
    }

    if price.Currency() != "USD" {
        t.Errorf("Expected USD, got %v", price.Currency())
    }
}

func TestNewPriceP(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Expected panic for invalid price")
        }
    }()

    NewPriceP(0, "USD")
}