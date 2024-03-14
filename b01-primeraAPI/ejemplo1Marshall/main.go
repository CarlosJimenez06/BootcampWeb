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
	p := product{
		Name:      "Laptop",
		Price:     1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(p)					-> Print the struct in a unreadable format for JSON

	//Print - convert - DataFormated
	fmt.Println(string(jsonData)) //-> Print the struct in a readable format for JSON
}
