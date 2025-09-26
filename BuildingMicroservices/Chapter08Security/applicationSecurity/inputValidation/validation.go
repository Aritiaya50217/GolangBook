package inputValidation

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Request struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"email"`
	URL   string `json:"url" validate:"url"`
}

var validate = validator.New()

func Handler(w http.ResponseWriter, r *http.Request) {
	request := Request{}
	if err := json.NewEncoder(w).Encode(&request); err != nil {
		http.Error(w, "invalid request object", http.StatusBadRequest)
		return
	}
	if err := validate.Struct(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
