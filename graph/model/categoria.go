package model

type Categoria struct {
	ID        string  `json:"id"`
	Nome      string  `json:"nome"`
	Descricao *string `json:"descricao,omitempty"`
}
