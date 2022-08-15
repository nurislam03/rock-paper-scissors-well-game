package api

import (
	"github.com/go-chi/chi"
	"net/http"
)

// API ...
type API struct {
	router chi.Router
}

// NewAPI ...
func NewAPI() *API {
	api := &API{
		router: chi.NewRouter(),
	}

	api.RegisterRoutes()
	return api
}

//GetRouter ...
func (api *API) GetRouter() http.Handler {
	return api.router
}

//RegisterRoutes ...
func (api *API) RegisterRoutes() {
	r := chi.NewRouter()
	r.Mount("/rpsw", api.rpswGameHandler())
	api.router = r
}

// rpswGameHandler is responsible for handling all rock-paper-scissors-well-game related endpoints
func (api *API) rpswGameHandler() http.Handler {
	r := chi.NewRouter()
	r.Post("/play", api.playRPSW)
	return r
}
