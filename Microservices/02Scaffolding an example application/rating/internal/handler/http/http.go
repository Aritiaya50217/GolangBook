package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"movieexample.com/rating/internal/controller/rating"
	"movieexample.com/rating/pkg/model"
)

// Handler defines a rating service controller
type Handler struct {
	ctrl *rating.Controller
}

// New creates a new rating service HTTP handler
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// Handle handler PUT and GET /rating request
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	recordId := model.RecordId(r.FormValue("id"))
	if recordId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	recordType := model.RecordType(r.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(r.Context(), recordId, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	case http.MethodPost:
		userId := model.UserID(r.FormValue("userId"))
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err := h.ctrl.PutRating(r.Context(), recordId, recordType, &model.Rating{
			UserID: userId,
			Value:  model.RatingValue(v),
		}); err != nil {
			log.Printf("Repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
