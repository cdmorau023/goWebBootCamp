package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Greetings struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
}

//Creando mi primer router o servidor web simple con chi

func main() {

	router := chi.NewRouter()

	//Vamos a crear un endpoint llamado /greetings. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hello + nombre + apellido”
	router.Post("/greetings", CreateGreetingsHandler)

	http.ListenAndServe(":8080", router)

	// Ya teniendo definido mi router, puedo agregarle endpoints
	//o rutas para ello debemos agregar al router distintos handlers

}

func CreateGreetingsHandler(writer http.ResponseWriter, request *http.Request) {
	var greeting Greetings
	err := json.NewDecoder(request.Body).Decode(&greeting)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Bad request"))
		return
	}
	response := fmt.Sprintf("Hello %s %s", greeting.Nombre, greeting.Apellido)

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, response)

}
