package boxeadores

import "github.com/ValentinCoccimiglio/API-boxeadores/internal/domain"

type BoxeadorService struct {
	repo BoxeadorRepository
}

func NewBoxeadorService(repo BoxeadorRepository) *BoxeadorService {
	return &BoxeadorService{repo: repo}
}

func (s *BoxeadorService) GetBoxeadores() ([]domain.Boxeador, error) {
	return s.repo.GetBoxeadores()
}
