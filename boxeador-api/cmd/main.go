package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tuusuario/boxeador-api/cmd/api/handler"
)

func main() {
	router := mux.NewRouter()

	// Define las rutas
	router.HandleFunc("/boxeadores", handler.GetBoxeadores).Methods("GET")

	// Inicia el servidor
	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
