package router

import (
	"github.com/andreastihor/stockbit/pkg/movie"
	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router, movieHandler movie.HandlerInterface) {
	clientRoute := router.PathPrefix("/api/").Subrouter()
	clientRoute.HandleFunc("/rest/movies", movieHandler.HandleGetMovie).Methods("GET")
}
