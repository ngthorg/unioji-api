package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/unioji/unioji-api/graph/model"
)

// TodoRepo struct
type TodoRepo struct {
	DB *pg.DB
}

// GetTodoByID is a method in TodoRepo
func (r *TodoRepo) GetTodoByID(id string) (*model.Todo, error) {
	var todo model.Todo
	// err := r.DB.Model(&todo).
	// 	Relation("User").
	// 	Where("todo.id = ?", id).
	// 	First()
	todo = model.Todo{ID: id}

	err := r.DB.Select(todo)

	fmt.Println(todo)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}
