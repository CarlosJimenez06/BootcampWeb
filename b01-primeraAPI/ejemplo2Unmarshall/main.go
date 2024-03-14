package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	Name      string
	Price     int
	Published bool
}

func main() {
	// Data en formato JSON
	jsonData := `{"Name": "Laptop", "Price": 1500, "Published": true}`

	// Variable en la que se almacenará el JSON deserializado
	var p product

	// Deserializar el JSON
	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		fmt.Println(err)
	}

	//
	fmt.Println(p)

}
