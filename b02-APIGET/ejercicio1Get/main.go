package main

import (
	//"encoding/json"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	//"time"

	"github.com/go-chi/chi/v5"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func main() {

	// Opennig the file from the local directory
	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	// Reading the file
	//reader := bufio.NewReader(file)

	// Var to store the file content
	var jsonDataBytes []Product
	err = json.NewDecoder(file).Decode(&jsonDataBytes)
	if err != nil {
		fmt.Println("Error al decodificar el JSON: ", err.Error())
		return
	}

	/*
		// Adding the file content to the slice
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				break
			}

			jsonDataBytes = append(jsonDataBytes, line...)
		}
	*/
	//fmt.Println(string(jsonDataBytes[0:100]))

	/*
	*	CREATION OF ROUTER
	 */

	// Creation of router
	router := chi.NewRouter()

	// Creation of endpoint GET "products" [HANDLERS]
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("pong"))
	})

	// Get all products
	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonDataBytes)
	})

	// Get product by id
	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		product := findProductById(id, jsonDataBytes)
		if product == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(product)
	})

	// Get products more expensive than a price
	router.Get("/products/search/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		priceGt, err := strconv.Atoi(r.URL.Query().Get("priceGt"))
		fmt.Println(priceGt)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var productsFiltered []Product
		for _, product := range jsonDataBytes {
			if product.Quantity > priceGt {
				productsFiltered = append(productsFiltered, product)
			}
		}
		//product := findProductByPrice(priceGt, jsonDataBytes)

		if productsFiltered == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(productsFiltered)
	})

	// Starting the server
	http.ListenAndServe(":8080", router)

}

func findProductById(id int, products []Product) *Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func findProductByPrice(priceGt int, products []Product) *Product {
	for _, product := range products {
		if product.Quantity > priceGt {
			return &product
		}
	}
	return nil
}
