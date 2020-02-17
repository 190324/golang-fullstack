// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	"context"

	models_gen "example.com/gqlgen/app/graphql/api/models"
)

func (r *queryResolver) User(ctx context.Context) (*models_gen.User, error) {
	// panic(fmt.Errorf("not implemented"))
	user := &models_gen.User{
		Name: "text2",
		ID:   "2",
	}
	return user, nil
}
