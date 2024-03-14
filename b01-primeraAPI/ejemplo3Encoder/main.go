package main

import (
	"encoding/json"
	"os"
	//"fmt"
)

type MyData struct {
	ProductID string
	Price     float64
}

func main() {
	// Create some data
	data := MyData{
		ProductID: "123",
		Price:     100.0,
	}

	// Generates a stream
	myEncoder := json.NewEncoder(os.Stdout)

	// Encode the data and write it to the stream
	myEncoder.Encode(data)
}
