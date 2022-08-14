package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

// API ...
type API struct {
	router chi.Router
}

//GetRouter ...
func (api *API) GetRouter() http.Handler {
	return api.router
}

//RegisterRoutes ...
func (api *API) RegisterRoutes() {
	r := chi.NewRouter()

	api.router = r
}
