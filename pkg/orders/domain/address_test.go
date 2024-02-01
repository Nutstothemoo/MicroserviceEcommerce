package orders

import (
    "testing"
)

func TestNewAdress(t *testing.T) {
    _, err := NewAddress("", "street", "city", "postcode", "country")
    if err == nil || err.Error() != "name can not be empty" {
        t.Errorf("Expected 'name can not be empty' error, got %v", err)
    }

    _, err = NewAddress("name", "", "city", "postcode", "country")
    if err == nil || err.Error() != "street can not be empty" {
        t.Errorf("Expected 'street can not be empty' error, got %v", err)
    }

    _, err = NewAddress("name", "street", "", "postcode", "country")
    if err == nil || err.Error() != "city can not be empty" {
        t.Errorf("Expected 'city can not be empty' error, got %v", err)
    }

    _, err = NewAddress("name", "street", "city", "", "country")
    if err == nil || err.Error() != "postcode can not be empty" {
        t.Errorf("Expected 'postcode can not be empty' error, got %v", err)
    }

    _, err = NewAddress("name", "street", "city", "postcode", "")
    if err == nil || err.Error() != "country can not be empty" {
        t.Errorf("Expected 'country can not be empty' error, got %v", err)
    }

    ad, err := NewAddress("name", "street", "city", "postcode", "country")
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if ad.Name() != "name" {
        t.Errorf("Expected name to be 'name', got %v", ad.Name())
    }

    if ad.Street() != "street" {
        t.Errorf("Expected street to be 'street', got %v", ad.Street())
    }

    if ad.City() != "city" {
        t.Errorf("Expected city to be 'city', got %v", ad.City())
    }

    if ad.Postcode() != "postcode" {
        t.Errorf("Expected postcode to be 'postcode', got %v", ad.Postcode())
    }

    if ad.Country() != "country" {
        t.Errorf("Expected country to be 'country', got %v", ad.Country())
    }
}