package resolvers

import (
	"context"
	"io/ioutil"

	models_gen "example.com/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql"
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

func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (*models_gen.File, error) {
	// panic("not implemented")

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
