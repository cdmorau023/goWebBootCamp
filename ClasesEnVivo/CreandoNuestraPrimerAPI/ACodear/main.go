package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

//Creando mi primer router o servidor web simple con chi

func main() {

	router := chi.NewRouter()
	router.Get("/hello-World", func(w http.ResponseWriter, r *http.Request) {
		// Escribir un mensaje de respuesta
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello World"}`))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

	// Ya teniendo definido mi router, puedo agregarle endpoints
	//o rutas para ello debemos agregar al router distintos handlers

}
