// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	"XinAPI/app/http/models"
	generated "XinAPI/build/gqlgen/admin"
	models_gen "XinAPI/build/gqlgen/admin/models"
	"context"
	"fmt"

	"github.com/gobuffalo/pop"
)

func (r *mutationResolver) SetProduct(ctx context.Context, input models_gen.InputSetProduct) (*models_gen.SetProductPayload, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProductImages(ctx context.Context, obj *models.Product) ([]*models_gen.ProductImage, error) {
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

func (r *queryResolver) Product(ctx context.Context, id string) (*models_gen.Product, error) {
	// panic(fmt.Errorf("not implemented"))

	// db, _ := pop.Connect("solomo")

	// product := models.Product{}
	// err := db.Eager("ProductImages").Find(&product, id)

	// if err != nil {
	// 	fmt.Printf("ddd %v", err)
	// 	return nil, nil
	// }

	desp := "ccc"
	product := &models_gen.Product{
		ID:   12356,
		Name: "c456",
		Desp: &desp,
	}

	return product, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
