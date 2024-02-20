package main

type BoxeadorRepository interface {
	GetBoxeadores() ([]Boxeador, error)
	// Agrega más métodos según sea necesario
}
