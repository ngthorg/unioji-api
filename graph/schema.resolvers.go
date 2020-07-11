package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/unioji/unioji-api/graph/connections"
	"github.com/unioji/unioji-api/graph/generated"
	"github.com/unioji/unioji-api/graph/model"
	"github.com/unioji/unioji-api/graph/relay"
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
	globalID := relay.FromGlobalID(id)
	if globalID == nil {
		return nil, errors.New("Unknown node id")
	}

	switch globalID.Type {
	case "User":
		return r.UserRepo.GetUserByID(globalID.ID)
	case "Todo":
		return r.TodoRepo.GetTodoByID(globalID.ID)
	default:
		return nil, errors.New("Unknown node type")
	}
}

func (r *queryResolver) Nodes(ctx context.Context, after *string, before *string, first *int, last *int, ids []string) (*model.NodeConnection, error) {
	nodes := []model.Node{}
	for _, id := range ids {
		globalID := relay.FromGlobalID(id)
		if globalID == nil {
			return nil, errors.New(fmt.Sprintf("Unknown node id: %v", id))
		}
		switch globalID.Type {
		case "User":
			node, _ := r.UserRepo.GetUserByID(globalID.ID)
			nodes = append(nodes, node)
			break
		case "Todo":
			node, _ := r.TodoRepo.GetTodoByID(globalID.ID)
			nodes = append(nodes, node)
			break
		default:
			return nil, errors.New("Unknown node type")
		}
	}
	return connections.ConnectionFromSliceNodes(nodes, relay.ConnectionArgs{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}, relay.SliceMetaInfo{
		Length: len(nodes),
	})
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
	todos := []*model.Todo{
		{
			ID:        "1",
			Text:      "Test 1",
			Completed: true,
		},
		{
			ID:        "2",
			Text:      "Demo",
			Completed: false,
		},
		{
			ID:        "3",
			Text:      "Test 3",
			Completed: false,
		},
		{
			ID:        "4",
			Text:      "Test 4",
			Completed: false,
		},
		{
			ID:        "5",
			Text:      "Test 5",
			Completed: false,
		},
		{
			ID:        "6",
			Text:      "Test 6",
			Completed: false,
		},
		{
			ID:        "7",
			Text:      "Test 7",
			Completed: false,
		},
		{
			ID:        "8",
			Text:      "Test 8",
			Completed: false,
		},
		{
			ID:        "9",
			Text:      "Test 9",
			Completed: false,
		},
		{
			ID:        "10",
			Text:      "Test 10",
			Completed: false,
		},
	}

	return connections.ConnectionFromSliceTodos(todos, relay.ConnectionArgs{
		Before: before,
		After:  after,
		First:  first,
		Last:   last,
	}, relay.SliceMetaInfo{
		Length: 20,
	})
}

func (r *todoResolver) ID(ctx context.Context, obj *model.Todo) (string, error) {
	return relay.ToGlobalID("Todo", obj.ID), nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	if obj.User != nil {
		return obj.User, nil
	}

	return r.UserRepo.GetUserByID(obj.UserID)
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return relay.ToGlobalID("User", obj.ID), nil
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
