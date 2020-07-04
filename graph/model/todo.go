package model

// Todo type
type Todo struct {
	tableName struct{} `pg:"todos"`
	ID        string   `json:"id"`
	Text      string   `json:"text"`
	Completed bool     `json:"completed"`
	UserID    string   `json:"user_id" pg:"user_id"`
	User      *User    `json:"user" pg:"fk:user_id"`
}

// IsNode is a method in Todo
func (Todo) IsNode() {}

// IsSearchResult is a method in Todo
func (Todo) IsSearchResult() {}

// NewTodo type
type NewTodo struct {
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}
