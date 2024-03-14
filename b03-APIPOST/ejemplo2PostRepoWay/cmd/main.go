package main

import (
	"ejemplo2PostRepoWay/internal/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	handler := handlers.NewDefaultTask(nil, 0)

	router := chi.NewRouter()

	router.Post("/tasks", handler.Create())

	// Start the server
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
		return
	}
}
