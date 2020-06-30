package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/unioji/unioji-api/graph/model"
)

// TodoRepo struct
type TodoRepo struct {
	DB *pg.DB
}

// GetTodoByID is a method in TodoRepo
func (m *TodoRepo) GetTodoByID(id string) (*model.Todo, error) {
	todo := &model.Todo{ID: id}

	err := m.DB.Select(todo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// GetTodos is a method in TodoRepo
func (m *TodoRepo) GetTodos() ([]model.Todo, error) {
	var todos []model.Todo
	err := m.DB.Model(&todos).Select()

	if err != nil {
		return nil, err
	}

	return todos, nil
}
