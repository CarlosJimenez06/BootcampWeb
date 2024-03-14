package handlers

import (
	"encoding/json"
	"net/http"
)

// Product es un struct que contiene la informacion de un producto
type Product struct {
	Id       int
	Name     string
	Type     string
	Quantity int
	Price    float64
}

// ControllerProducts es un struct que contiene el STORAGE de productos
type ControllerProducts struct {
	storage map[int]*Product
}

// Aquí procesaremos la información que va en el body de la REQUEST
type RequestBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type ResponseBodyProduct struct {
	Message string   `json:"message"`
	Data    *Product `json:"data"`
	Error   bool     `json:"error"`
}

// Maneja el REQUEST y GUARDA el Product
func (c *ControllerProducts) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// - REQUEST
		var reqBody RequestBodyProduct
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			code := http.StatusBadRequest
			body := &ResponseBodyProduct{
				Message: "Bad Request",
				Data:    nil,
				Error:   true,
			}

			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)
		}

		// - PROCESS
		// Deserialization
		pr := &Product{
			Id:       len(c.storage) + 1,
			Name:     reqBody.Name,
			Type:     reqBody.Type,
			Quantity: reqBody.Quantity,
			Price:    reqBody.Price,
		}
		// Save the product in the storage
		c.storage[pr.Id] = pr

		// - RESPONSE
		code := http.StatusCreated
		body := &ResponseBodyProduct{
			Message: "Product created successfully",
			Data:    pr,
			Error:   false,
		}

		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
}

/*
// Esta es una representación de la RESPUESTA que va a tener la petición
type ResponseBodyProduct struct {
	Message string `json:"message"`
	Data *struct {
		Id int `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		Quantity int `json:"quantity"`
		Price float64 `json:"price"`
	} `json:"data"`
	Error bool `json:"error"`
}*/
