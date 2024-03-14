package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/go-chi/chi/v5"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	// Verificar que el método de la solicitud sea POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar el cuerpo JSON de la solicitud en una estructura Person
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Componer el mensaje de saludo
	message := fmt.Sprintf("Hello %s %s", person.FirstName, person.LastName)

	// Responder con el mensaje de saludo
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, message)
}

func main() {
	http.HandleFunc("/greetings", greetingsHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
