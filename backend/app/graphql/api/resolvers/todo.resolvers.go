// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	"context"
	"fmt"
	"io/ioutil"

	generated "example.com/gqlgen/app/graphql/api"
	models_gen "example.com/gqlgen/app/graphql/api/models"
	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input models_gen.NewTodo) (*models_gen.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*models_gen.File, error) {
	// panic(fmt.Errorf("not implemented"))
	content, err := ioutil.ReadAll(file.File)
	if err != nil {
		return nil, err
	}

	f := &models_gen.File{
		ID:      1,
		Name:    file.Filename,
		Content: string(content),
	}

	return f, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models_gen.Todo, error) {
	// panic(fmt.Errorf("not implemented"))
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

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
