package handler

import (
	"apiboxeadores/pkg/domain"
	"encoding/json"
	"net/http"

	"github.com/tuusuario/boxeador-api/pkg/service"
)

var boxeadorService *service.BoxeadorService

func init() {
	// Aquí puedes inicializar el repositorio y el servicio si es necesario
	// Ejemplo: boxeadorService = service.NewBoxeadorService(repository.NewTuRepositorio())
}

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
