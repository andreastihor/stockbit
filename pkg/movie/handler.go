package movie

import (
	"net/http"

	"github.com/andreastihor/stockbit/helper"
	"github.com/andreastihor/stockbit/models"
	"github.com/gorilla/schema"
)

type Handler struct {
	Service ServiceInterface
}

type HandlerInterface interface {
	HandleGetMovie(w http.ResponseWriter, r *http.Request)
}

// NewHandler will initialize the user endpoint
func NewHandler(service ServiceInterface) HandlerInterface {
	movieHandler := &Handler{
		Service: service,
	}

	return movieHandler
}

// HandleGetMovie
func (h *Handler) HandleGetMovie(w http.ResponseWriter, r *http.Request) {
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
