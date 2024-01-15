package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Definición de una estructura para almacenar datos JSON decodificados
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Datos JSON de ejemplo
	jsonData := `{"name":"Alice","age":30}`

	// Crear un decoder basado en una cadena de JSON
	decoder := json.NewDecoder(strings.NewReader(jsonData))

	// Crear una variable para almacenar los datos decodificados
	var p Person

	// Decodificar los datos JSON en la estructura Person
	err := decoder.Decode(&p)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return
	}

	// Imprimir los datos decodificados
	fmt.Println("Nombre:", p.Name)
	fmt.Println("Edad:", p.Age)
}
