package resolvers

import (
	"context"

	generated "example.com/gqlgen/app/graphql/admin"
	models_gen "example.com/gqlgen/app/graphql/admin/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models_gen.NewTodo) (*models_gen.Todo, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models_gen.Todo, error) {
	// panic("not implemented")
	return nil, nil
}
