package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type product struct {
	ID        int
	Name      string
	Price     float64
	Available bool
}

func jsonMarshalExample() {
	p := product{
		ID:        1,
		Name:      "Laptop",
		Price:     1500,
		Available: true,
	}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
func jsonUnmarshalExample() {
	jsonData := `{"id":1,"name":"Laptop","price":1500,"available":true}`
	var p product
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}

//Ejemplo función marshal

func main() {
	jsonMarshalExample()
	jsonUnmarshalExample()

}
