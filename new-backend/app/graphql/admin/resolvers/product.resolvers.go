// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	generated "XinAPI/app/graphql/admin"
	models_gen "XinAPI/app/graphql/admin/models"
	"context"
	"fmt"
)

func (r *mutationResolver) SetProduct(ctx context.Context, input models_gen.InputSetProduct) (*models_gen.SetProductPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (*models_gen.Product, error) {
	// panic(fmt.Errorf("not implemented"))

	desp := "描述中"

	return &models_gen.Product{
		ID:   "1",
		Name: "產品一號",
		Desp: &desp,
	}, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
