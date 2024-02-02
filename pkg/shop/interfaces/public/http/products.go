package http

import (
	"net/http"
	"github.com/go-chi/render"
	products_domain "microservice/pkg/orders/domain"
	common_http "microservice/pkg/common/http"
	"microservice/pkg/shop/domain"
	"microservice/pkg/common/price"
	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux, productsReadModel productsReadModel) {
	resource := productsResource{productsReadModel}
	router.Get("/products", resource.GetAll)
	router.Get("/products/{id}", resource.GetByID)
}

type productsReadModel interface {
	AllProducts() ([]products.Product, error)
	ByID(products_domain.ProductID) (products.Product, error)
}
type productsResource struct {
	readModel productsReadModel
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

func (p productsResource) GetByID(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")
	product, err := p.readModel.ByID(products_domain.ProductID(productID))
	if err != nil {
			_ = render.Render(w, r, common_http.ErrInternal(err))
			return
	}

	view := productView{
			ID:          string(product.ID()),
			Name:        product.Name(),
			Description: product.Description(),
			Price:       priceViewFromPrice(product.Price()),
	}

	render.Respond(w, r, view)
}

