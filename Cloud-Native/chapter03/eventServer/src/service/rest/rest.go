package rest

import (
	"Cloud-Native/chapter02/event/src/lib/persistence"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeAPI(endpoint, tlsEndpoint string, databaseHandler persistence.DatabaseHandler) (chan error, chan error) {
	handler := NewEventHandler(databaseHandler)
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)
	httpErrChan := make(chan error)
	httptlsErrChan := make(chan error)

	go func() {
		httptlsErrChan <- http.ListenAndServeTLS(tlsEndpoint, "cert.pem", "key.pem", r)
	}()

	go func() {
		httpErrChan <- http.ListenAndServe(endpoint, r)
	}()

	return httpErrChan, httptlsErrChan
}
