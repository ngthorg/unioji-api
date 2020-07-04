package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/unioji/unioji-api/graph/generated"
	"github.com/unioji/unioji-api/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:        "23423",
		Text:      input.Text,
		Completed: input.Completed,
	}
	return todo, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*model.User, error) {
	// default query user id 1
	// sau này sẽ dựa token
	return r.UserRepo.GetUserByID("1")
}

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	todo, err := r.TodoRepo.GetTodoByID(id)

	if err != nil {
		return nil, err
	}

	return *todo, nil
}

func (r *queryResolver) Search(ctx context.Context, text string) ([]model.SearchResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]model.User, error) {
	return r.UserRepo.GetUsers()
}

func (r *queryResolver) Todos(ctx context.Context) ([]model.Todo, error) {
	return r.TodoRepo.GetTodos()
}

func (r *queryResolver) TodosConnection(ctx context.Context, after *string, before *string, first *int, last *int, orderBy *model.TodoOrderBy) (*model.TodoConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj.User != nil {
		return obj.User, nil
	}

	return r.UserRepo.GetUserByID(obj.UserID)
}

func (r *userResolver) Friends(ctx context.Context, obj *model.User) ([]model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) TodosConnection(ctx context.Context, obj *model.User, after *string, before *string, first *int, last *int, orderBy *model.TodoOrderBy) (*model.TodoConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
