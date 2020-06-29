package graph

import "github.com/unioji/unioji-api/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct
type Resolver struct {
	TodoRepo postgres.TodoRepo
}
