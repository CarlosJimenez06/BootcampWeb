package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	// CREATING HANDLERS
	// Create a new endpoint GET "hello/world"
	router.Get("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})

	//rt := chi.NewRouter()

	http.ListenAndServe(":8080", router)
}
