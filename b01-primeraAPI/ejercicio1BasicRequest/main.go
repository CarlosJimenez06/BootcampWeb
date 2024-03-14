package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	// Creating a new router
	rt := chi.NewRouter()

	// Creating a new endpoint GET "hello/world" [HANDLERS]
	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("pong"))
	})

	// Starting the server
	http.ListenAndServe(":8080", rt)
}
