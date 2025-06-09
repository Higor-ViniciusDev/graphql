package model

type Curso struct {
	ID        string     `json:"id"`
	Nome      string     `json:"nome"`
	Descricao *string    `json:"descricao,omitempty"`
	Categoria *Categoria `json:"categoria"`
}
