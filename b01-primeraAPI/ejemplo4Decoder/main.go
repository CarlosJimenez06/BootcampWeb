package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var jsonStream = `{"ProductID": "123", "Price": 100.0}
{"ProductID": "456", "Price": 200.0}
{"ProductID": "789", "Price": 300.0}`

var myStreaming = strings.NewReader(jsonStream)
var myDecoder = json.NewDecoder(myStreaming)

type MyData struct {
	ProductID string
	Price     float64
}

func main() {
	fmt.Println(myDecoder)
	for {
		var data MyData
		err := myDecoder.Decode(&data)
		if err != nil {
			break
		}
		fmt.Println(data)
	}

	//println("--------------------\n")
	//println(jsonStream)
}
