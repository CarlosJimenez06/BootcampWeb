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
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Exp         string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {

	// Openning the file from the local directory
	file, err := os.Open("products.json")
	// Checking for errors
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	// Var to store the file content
	var dataProducts = []Product{}
	// Decoding the file
	err = json.NewDecoder(file).Decode(&dataProducts)
	// Checking for errors
	if err != nil {
		fmt.Println("Error al decodificar el JSON: ", err.Error())
		return
	}

	// Creating the router
	rt := chi.NewRouter()

	// Endpoint GET "/ping"
	rt.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	// Endpoint GET "/products" to get all the products
	rt.Get("/products2", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataProducts)
	})

	// Endpoint GET "/products/{id}" to get a product by its ID
	rt.Get("/products2/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Error al convertir el ID a entero: ", err.Error())
		}
		product := findProductByID(id, dataProducts)
		if product == nil {
			http.Error(w, "Producto no encontrado", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(product)
	})

	// Endpoint GET "/products/search" to get a product more than a price
	rt.Get("/products2/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		priceStr := r.URL.Query().Get("price")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Println("Error al convertir el precio a float: ", err.Error())
		}
		var products []Product
		for _, product := range dataProducts {
			if product.Price > price {
				products = append(products, product)
			}
		}
		json.NewEncoder(w).Encode(products)
	})

	// Testing
	rt.Get("/movies", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		titleLike := r.URL.Query().Get("title_like")
		awards := r.URL.Query().Get("awards")

		w.Write([]byte(
			"title_like: " + titleLike + "awards: " + awards,
		))
	})

	// Starting the server

	http.ListenAndServe(":8080", rt)
}

func findProductByID(id int, products []Product) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
