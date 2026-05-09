package transport

import (
	"building_a_microservice_using_DDD/recommendation/internal/recommendation"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMux(recHandler recommendation.Handler) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/recommendation", recHandler.GetRecommendation).Methods(http.MethodGet)
	return m
}
