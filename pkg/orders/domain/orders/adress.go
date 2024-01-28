package orders

import "errors"

type Adress struct {
	name string
	street string
	city string
	postcode string
	country string
}


func NewAdress(name, street, city, postcode, country string) (*Adress, error) {
	if name == "" {
		return nil, errors.New("name can not be empty")
	}
	if street == "" {
		return nil, errors.New("street can not be empty")
	}
	if city == "" {
		return nil, errors.New("city can not be empty")
	}
	if postcode == "" {
		return nil, errors.New("postcode can not be empty")
	}
	if country == "" {
		return nil, errors.New("country can not be empty")
	}
	return &Adress{name, street, city, postcode, country}, nil
}

func (a *Adress) Name() string {
	return a.name
}
func (a *Adress) Street() string {
	return a.street
}
func (a *Adress) City() string {
	return a.city
}
func (a *Adress) Postcode() string {
	return a.postcode
}
func (a *Adress) Country() string {
	return a.country
}
