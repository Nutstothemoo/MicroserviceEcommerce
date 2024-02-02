package application

import (
	"microservice/pkg/common/price"
	"github.com/pkg/errors"
	"microservice/pkg/shop/domain"
)

type productReadModel interface {
	AllProducts() ([]products.Product, error)
	AddProduct(cmd AddProductCommand) error
	ProductById(id string) (products.Product, error)
	// CreateProduct(cmd CreateProductCommand) error
	// RemoveProduct(cmd RemoveProductCommand) error
	// UpdateProduct(cmd UpdateProductCommand) error
}

type productsService struct {
	repo products.Repository
	readModel productReadModel
}

func NewProductsService(repo products.Repository, readModel productReadModel) productsService {
	return productsService{
		repo: repo,
		readModel: readModel,
	}
}
func(s productsService) AllProducts() ([]products.Product, error) {
	return s.readModel.AllProducts()
}

func(s productsService) ProductById(id string) (products.Product, error) {
	return products.Product{}, nil
}

type AddProductCommand struct {
	ID 					string 
	Name 				string 
	PriceCents 	uint 
	Description string 
	Currency 		string 
}

func (s productsService)AddProduct(cmd AddProductCommand) error {

	price, err := price.NewPrice(cmd.PriceCents, cmd.Currency)
	if err != nil {
		return errors.Wrap(err, "Error creating price")
	}

	p, err := products.NewProduct(products.ID(cmd.ID), cmd.Name, price, cmd.Description)
	if err != nil {
		return errors.Wrap(err, "Error creating product")
	}
	if err:= s.repo.Save(p); err != nil {
		return errors.Wrap(err, "Error saving product")
	}
	return nil
}

// func()CreateProduct(cmd CreateProductCommand) error {
// 	return nil
// }

// func (s productsService)RemoveProduct(cmd RemoveProductCommand) error {
// 	return nil
// }

// func (s productsService)UpdateProduct(cmd UpdateProductCommand) error {	
// 	return nil
// }
