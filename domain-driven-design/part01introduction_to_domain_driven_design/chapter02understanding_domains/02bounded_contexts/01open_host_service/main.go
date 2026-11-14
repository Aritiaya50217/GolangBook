package openhostservice

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler interface {
	IsUserSubscriptionActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	m := mux.NewRouter()
	m.HandleFunc("/user/{userID}/subscription/active", func(w http.ResponseWriter, r *http.Request) {
		// check auth
		uID := mux.Vars(r)["userID"]
		if uID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		isActive := u.IsUserSubscriptionActive(r.Context(), uID)
		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(b)
	}).Methods(http.MethodGet)
}
