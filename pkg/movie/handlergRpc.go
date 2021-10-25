package movie

import (
	"context"
	"net/http"

	"github.com/andreastihor/stockbit/helper"
	"github.com/andreastihor/stockbit/models"
	"github.com/gorilla/schema"
)

type HandlerRPC struct {
	Service models.ServiceClient
}

type HandlerRPCInterface interface {
	GetMoviesRPC(w http.ResponseWriter, r *http.Request)
}

// NewService will initialize the service for  user endpoint
func NewHandlerRPC(s models.ServiceClient) HandlerRPCInterface {
	return &HandlerRPC{
		s,
	}
}

// GetMoviesRPC
func (h *HandlerRPC) GetMoviesRPC(w http.ResponseWriter, r *http.Request) {
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

	param := models.ParamsProto{
		Pagination: params.Pagination,
		Search:     params.Search,
	}

	data, err := h.Service.GetMovies(context.Background(), &param)
	if err != nil {
		helper.ResponseHandler(w, nil, err.Error(), 500)
		return
	}

	helper.ResponseHandler(w, data, "", 200)
}
