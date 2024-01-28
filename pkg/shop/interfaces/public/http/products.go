package http

import (
	"net/http"
	"github.com/go-chi/render"
	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux, productsReadModel products_domain.ReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
	router.Get("/products/{id}", resource.GetByID)
}

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
	ByID(products_domain.ProductID) (products.Product, error)
}

type productView struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       PriceView `json:"price"`

}

type PriceView struct {
	Cents    uint   `json:"cents"`
	Currency string `json:"currency"`
}

type productsResource struct {
	readModel productsReadModel
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.readModel.AllProducts()
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	view := []productView{}
	for _, product := range products {
		view = append(view, productView{
			ID:          string(product.ID()),
			Name:        product.Name(),
			Description: product.Description(),
			Price:       priceViewFromPrice(product.Price()),
		})
	}
	render.Respond(w, r, view)
}

func priceViewFromPrice(p price.Price) PriceView {
	return PriceView{
		Cents:    p.Cents(),
		Currency: p.Currency(),
	}
}

