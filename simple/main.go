package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.HandleFunc("/ishealthy", handleIsHealthy).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}

// MIDDLEWARE handler
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Url requested: %s", r.RequestURI)
		next.ServeHTTP(w, r)

		// all remaining handlers on the pipeline will execute before reaching this line
		log.Println("Request finished")
	})
}

// URL handler
func handleIsHealthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Returning 200 - Healthy")
	w.Write([]byte("Healthy"))
}
