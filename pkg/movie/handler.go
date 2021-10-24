package movie

import (
	"net/http"

	"github.com/andreastihor/stockbit/helper"
	"github.com/andreastihor/stockbit/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Handler struct {
	Service ServiceInterface
}

// NewHandler will initialize the user endpoint
func NewHandler(router *mux.Router, service ServiceInterface) {
	movieHandler := &Handler{
		Service: service,
	}

	clientrouter := router.PathPrefix("/api/").Subrouter()
	clientrouter.HandleFunc("/movies", movieHandler.GetMovies).Methods("GET")
}

// GetMovies
func (h *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	params := &models.Params{}
	err := r.ParseForm()
	if err != nil {
		helper.ResponseHandler(w, nil, err.Error(), 400)
		return
	}

	err = schema.NewDecoder().Decode(params, r.Form)
	if err != nil {
		helper.ResponseHandler(w, nil, err.Error(), 400)
		return
	}

	data, err := h.Service.GetMovies(params)
	if err != nil {
		helper.ResponseHandler(w, nil, err.Error(), 500)
		return
	}

	helper.ResponseHandler(w, data, "", 200)
}
