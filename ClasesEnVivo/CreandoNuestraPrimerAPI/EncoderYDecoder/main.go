package main

import (
	"encoding/json"
	"os"
)

type MyData struct {
	ProductID string
	Price     float64
}

func main() {
	// Crear un encoder
	myEncoder := json.NewEncoder(os.Stdout)
	data := MyData{
		ProductID: "1",
		Price:     1500,
	}
	//
	myEncoder.Encode(data)

}
