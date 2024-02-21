package service

import (
	"apiboxeadores/pkg/domain/repository"

	"github.com/tuusuario/boxeador-api/pkg/domain"
)

type BoxeadorService struct {
	repo repository.BoxeadorRepository
}

func NewBoxeadorService(repo repository.BoxeadorRepository) *BoxeadorService {
	return &BoxeadorService{repo: repo}
}

func (s *BoxeadorService) GetBoxeadores() ([]domain.Boxeador, error) {
	return s.repo.GetBoxeadores()
}
