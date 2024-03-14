package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Greeting struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {
	/*
		// Decoding the request body
		jsonStream := `{"first_name": "John", "last_name": "Doe"}`
		myNR := strings.NewReader(jsonStream)
		myDecoder := json.NewDecoder(myNR)
	*/
	jsonData := `{"first_name": "John", "last_name": "Doe"}`
	var p Greeting

	// Deserializar el JSON
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		fmt.Println(err)
	}

	// Creating a new router
	rt := chi.NewRouter()

	/* USING UNMARSHALL*/
	rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello " + p.FirstName + " " + p.LastName))
	})

	/*	USING DECODER
		// Creating a new endpoint POST "greetings"
		rt.Post("/greetings", func(w http.ResponseWriter, r *http.Request) {
			var greeting Greeting
			err := myDecoder.Decode(&greeting)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Bad request"))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte("Hello " + greeting.FirstName + " " + greeting.LastName))
		})
	*/

	// Starting the server
	http.ListenAndServe(":8080", rt)
}
