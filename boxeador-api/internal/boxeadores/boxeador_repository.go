package boxeadores

import "github.com/tuusuario/boxeador-api/internal/domain"

type BoxeadorRepository interface {
	GetBoxeadores() ([]domain.Boxeador, error)
}
