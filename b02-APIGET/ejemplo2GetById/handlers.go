package handlers

import (
	"encoding/json"
	"net/http"

	//"context"
	"github.com/go-chi/chi/v5"
)

type ControllerEmployee struct {
	// storage
	st map[string]string
}

// ResponseGetByIdEmployee is the response of the API Get By Id Employee
type ResponseGetByIdEmployee struct {
	Message string
	Data    *struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
	Error bool `json:"error"`
}

func (c *ControllerEmployee) ResponseGetByIdEmployee() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		id := chi.URLParam(r, "id")

		// Process
		// -> get employee by id
		employee, ok := c.st[id]
		if !ok {
			code := http.StatusNotFound
			body := &ResponseGetByIdEmployee{Message: "Employee not found", Data: nil, Error: true}

			w.WriteHeader(code)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
			return
		}

		// Response
		code := http.StatusOK
		body := &ResponseGetByIdEmployee{Message: "Employee found", Data: &struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{Id: id, Name: employee}, Error: false}

		w.WriteHeader(code)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
}
