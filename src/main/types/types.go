package types

type Materia struct {
	Id           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Nivel        int    `json:"nivel"`
	Especialidad string `json:"especialidad"`
}
