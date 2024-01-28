package http


import (
	"encoding/json"
	"net/http"
	"github.com/go-chi/render"
	"github.com/go-chi/chi"
)

func AddRoutes(router *chi.Mux , repo products_domain.MemoryRepository){
	resource := productsResource{repo}
	router.Get("/products", resource.GetAll)
	router.Get("/products/{id}", resource.GetByID)	
}

type productsResource struct {
	repo products_domain.Repository
}

type PriceView struct {
	Cents uint 				`json:"cents"`
	Currency string 	`json:"currency"`
}

func priceViewFromPrice(p price.Price) PriceView {
	return PriceView{
		Cents: p.Cents(),
		Currency: p.Currency(),
	}
}

func (p productsResource) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := p.repo.AllProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}


func (p productsResource) GetByID(w http.ResponseWriter, r *http.Request) {
	product, err := p.repo.ByID(products_domain.ProductID(r.URL.Query().Get("id")))
	if err != nil {
		_ = render.Render(w, r, common_http.ErrInternal(err))
		return
	}
	render.Respond(w, r, ProductView{
		string(product.ID()),
		product.Name(),
		product.Description(),
		priceViewFromPrice(product.Price()),
	})
}