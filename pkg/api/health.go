package api

import (
	"log"
	"net/http"
)

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(`{"status": "ok"}`))
	if err != nil {
		log.Printf("error writing response: %v", err)
	}
}
