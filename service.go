package main

type BoxeadorService struct {
	repo BoxeadorRepository
}

func NewBoxeadorService(repo BoxeadorRepository) *BoxeadorService {
	return &BoxeadorService{repo: repo}
}

func (s *BoxeadorService) GetBoxeadores() ([]Boxeador, error) {
	return s.repo.GetBoxeadores()
}
