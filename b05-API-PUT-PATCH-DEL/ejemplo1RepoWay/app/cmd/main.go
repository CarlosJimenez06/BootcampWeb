package main

import (
	"ejemplo1RepoWay/app/internal/handlers"
	"ejemplo1RepoWay/app/internal/repository"
	"ejemplo1RepoWay/app/internal/service"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	rp := repository.NewTaskMap(nil, 0)
	s := service.NewDefaultTask(rp)
	hd := handlers.NewDefaultTask(s)

	rt := chi.NewRouter()
	rt.Route("/tasks", func(rt chi.Router) {
		rt.Post("/", hd.Create())
		rt.Put("/{id}", hd.Update())
		rt.Delete("/{id}", hd.Delete())
		rt.Patch("/{id}", hd.UpdatePartial())
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		fmt.Println(err)
	}
}
