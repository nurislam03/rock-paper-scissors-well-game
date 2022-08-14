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
	r.Mount("/rpsw", api.rpswGameRouter())
	api.router = r
}

// rpswGameRouter ...
func (api *API) rpswGameRouter() http.Handler {
	r := chi.NewRouter()

	r.Post("/play", api.playRPSW)

	return r
}
