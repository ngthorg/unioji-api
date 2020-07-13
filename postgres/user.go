package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/unioji/unioji-api/graph/model"
)

// UserRepo struct
type UserRepo struct {
	DB *pg.DB
}

// GetUserByID is a method in UserRepo
func (m *UserRepo) GetUserByID(id string) (*model.User, error) {
	user := new(model.User)
	err := m.DB.Model(user).Where("id = ?", id).Relation("Todos").Select()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUsers is a method in UserRepo
func (m *UserRepo) GetUsers() ([]model.User, error) {
	users := new([]model.User)
	err := m.DB.Model(users).Relation("Todos").Select()

	if err != nil {
		return nil, err
	}

	return *users, nil
}

// SearchUsers in a method in UserRepo
func (m *UserRepo) SearchUsers(text string) ([]model.User, error) {
	users := new([]model.User)
	err := m.DB.Model(users).Relation("Todos").Where("name LIKE ?", "%"+text+"%").Select()
	if err != nil {
		return nil, err
	}
	return *users, nil
}
