package repository

import "github.com/tuusuario/boxeador-api/pkg/domain"

type BoxeadorRepository interface {
	GetBoxeadores() ([]domain.Boxeador, error)
	// Agrega más métodos según sea necesario
}
