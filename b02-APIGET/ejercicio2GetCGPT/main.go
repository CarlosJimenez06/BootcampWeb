package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	// Abrir el archivo JSON
	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	// Decodificar el JSON en una lista de productos
	var products []Product
	err = json.NewDecoder(file).Decode(&products)
	if err != nil {
		fmt.Println("Error al decodificar el JSON:", err)
		return
	}

	// Creación del router
	router := chi.NewRouter()

	// Endpoint GET "/ping"
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// Endpoint GET "/products" para obtener todos los productos
	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	})

	// Endpoint GET "/products/{id}" para obtener un producto por su ID
	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		product := findProductByID(id, products)
		if product == nil {
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(product)
	})

	// Endpoint GET "/products/search" para buscar productos por precio mayor a un valor dado
	router.Get("/products/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		priceGtStr := r.URL.Query().Get("priceGt")
		priceGt, err := strconv.ParseFloat(priceGtStr, 64)
		if err != nil {
			http.Error(w, "Parámetro 'priceGt' inválido", http.StatusBadRequest)
			return
		}

		var filteredProducts []Product
		for _, p := range products {
			if p.Price > priceGt {
				filteredProducts = append(filteredProducts, p)
			}
		}

		json.NewEncoder(w).Encode(filteredProducts)
	})

	// Iniciar el servidor
	http.ListenAndServe(":8080", router)
}

// Función para buscar un producto por su ID
func findProductByID(id int, products []Product) *Product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
