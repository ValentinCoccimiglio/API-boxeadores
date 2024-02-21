package domain

type Boxeador struct {
	ID       string `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Peso     string `json:"peso"`
	Edad     int    `json:"edad"`
}
