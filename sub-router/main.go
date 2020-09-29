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

	subRouter := router.PathPrefix("/sub").Subrouter()
	subRouter.Use(subMiddleware)
	subRouter.HandleFunc("/a", handleSubHealthyA).Methods(http.MethodGet)
	subRouter.HandleFunc("/b", handleSubHealthyB).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}

// ROUTER

// pipeline: loggingMiddleware -> handleIsHealthy

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Url requested: %s", r.RequestURI)
		next.ServeHTTP(w, r)
		log.Println("Request finished")
	})
}

func handleIsHealthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Returning 200 - Healthy")
	w.Write([]byte("Healthy"))
}

// SUB-ROUTER

// pipeline: loggingMiddleware -> subMiddleware -> [ handleIsHealthyA | handleIsHealthyB ]

func subMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Another middleware")
		next.ServeHTTP(w, r)
	})
}

func handleSubHealthyA(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Returning 200 - Healthy a")
	w.Write([]byte("Healthy a"))
}

func handleSubHealthyB(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Returning 200 - Healthy b")
	w.Write([]byte("Healthy b"))
}
