// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	generated "XinAPI/app/graphql/admin"
	models_gen "XinAPI/app/graphql/admin/models"
	"XinAPI/app/http/models"
	"context"
	"fmt"

	"github.com/gobuffalo/pop"
)

func (r *mutationResolver) SetProduct(ctx context.Context, input models_gen.InputSetProduct) (*models_gen.SetProductPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) ProductImages(ctx context.Context, obj *models.Product) ([]*models_gen.ProductImage, error) {
	// panic(fmt.Errorf("not implemented"))
	db, _ := pop.Connect("solomo")

	productImages := []models.ProductImage{}
	query := db.Where("product_id = ?", obj.ID)
	err := query.All(&productImages)

	output := []*models_gen.ProductImage{}
	for _, row := range productImages {
		output = append(output, &models_gen.ProductImage{
			ID:     row.ID,
			Image:  row.Image,
			Seq:    row.Seq,
			IsMain: row.IsMain,
		})
	}

	fmt.Println(err)

	return output, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*models.Product, error) {
	// panic(fmt.Errorf("not implemented"))

	db, _ := pop.Connect("solomo")

	product := models.Product{}
	err := db.Eager("ProductImages").Find(&product, id)

	if err != nil {
		fmt.Printf("ddd %v", err)
		return nil, nil
	}

	return &product, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Product() generated.ProductResolver   { return &productResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
