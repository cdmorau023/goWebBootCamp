package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Product representa la estructura de un producto.
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
	sliceProductos, err := loadProducts("products.json")
	if err != nil {
		fmt.Printf("Error al cargar los productos: %v", err)
		return
	}

	router := chi.NewRouter()

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		productJSON, err := json.Marshal(sliceProductos)
		if err != nil {
			w.Write([]byte("Error al serializar los productos"))
			return
		}
		w.Write(productJSON)
	})

	// Nuevo handler para obtener un producto por ID
	router.Get("/products/{id}", func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "id")
		productID, err := strconv.Atoi(idParam)
		if err != nil {
			w.Write([]byte("ID de producto no válido"))
			return
		}

		// Buscar el producto en la slice por su ID
		var foundProduct *Product
		for _, p := range sliceProductos {
			if p.ID == productID {
				foundProduct = &p
				break
			}
		}

		// Verificar si el producto fue encontrado
		if foundProduct == nil {
			w.Write([]byte("Producto no encontrado"))
			return
		}

		// Convertir el producto encontrado a JSON y enviar la respuesta
		productJSON, err := json.Marshal(foundProduct)
		if err != nil {
			w.Write([]byte("Error al serializar el producto"))
			return
		}
		w.Write(productJSON)
	})

	http.ListenAndServe(":8080", router)
}

// Función que dado el nombre de un archivo json retorna un slice con todos los productos
func loadProducts(filename string) ([]Product, error) {
	archivo, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(archivo)
	if err != nil {
		return nil, err
	}
	var products []Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}
