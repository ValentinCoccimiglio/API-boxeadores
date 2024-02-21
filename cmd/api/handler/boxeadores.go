package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ValentinCoccimiglio/API-boxeadores/internal/domain"
)

func GetBoxeadores(w http.ResponseWriter, r *http.Request) {
	// Lógica para obtener boxeadores desde el servicio
	// Puedes leer desde un repositorio en memoria o desde una base de datos

	boxeadores := []domain.Boxeador{
		{ID: "1", Nombre: "Mike Tyson", Peso: "Pesado", Edad: 55},
		{ID: "2", Nombre: "Muhammad Ali", Peso: "Pesado", Edad: 74},
		// Agrega más boxeadores según sea necesario
	}

	// Convierte los boxeadores a formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boxeadores)
}
