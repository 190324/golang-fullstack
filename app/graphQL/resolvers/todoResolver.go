package resolvers

import (
	"context"

	models_gen "example.com/gqlgen/graphql"
)

func (r *queryResolver) Todos(ctx context.Context) ([]*models_gen.Todo, error) {
	// panic("not implemented")
	todo := []*models_gen.Todo{
		{
			Text: "text",
			ID:   "1",
			Done: false,
			User: nil,
		},
		{
			Text: "text2",
			ID:   "2",
			Done: true,
			User: nil,
		},
	}
	return todo, nil
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input models_gen.NewTodo) (*models_gen.Todo, error) {
	panic("not implemented")
}
