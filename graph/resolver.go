package graph

import "github.com/Higor-ViniciusDev/graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CategoriaDB *database.Categoria
	CursoDB     *database.Curso
}
