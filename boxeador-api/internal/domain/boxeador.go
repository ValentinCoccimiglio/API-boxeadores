package domain

type Boxeador struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Peso   string `json:"peso"`
	Edad   int    `json:"edad"`
	// Agrega más campos según sea necesario
}
