// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	generated "XinAPI/app/graphql/api"
	models_gen "XinAPI/app/graphql/api/models"
	"XinAPI/app/mail"
	"context"
)

func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (*string, error) {
	// panic(fmt.Errorf("not implemented"))

	var str = "success"

	mail.ForgotPassword(email)

	return &str, nil
}

func (r *queryResolver) Me(ctx context.Context) (*models_gen.Member, error) {
	// panic(fmt.Errorf("not implemented"))

	return &models_gen.Member{
		ID:   "123",
		Name: "hello world",
	}, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
