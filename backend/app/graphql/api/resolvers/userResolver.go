package resolvers

import (
	"context"

	models_gen "example.com/gqlgen/app/graphql/api/models"
)

func (r *queryResolver) User(ctx context.Context) (*models_gen.User, error) {
	// panic("not implemented")
	user := &models_gen.User{
		Name: "text2",
		ID:   "2",
	}
	return user, nil
}
