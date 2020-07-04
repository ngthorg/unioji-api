package model

// User type
type User struct {
	tableName struct{} `pg:"users"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Todos     []*Todo  `json:"todos"`
}

// IsNode is a method in User
func (User) IsNode() {}
