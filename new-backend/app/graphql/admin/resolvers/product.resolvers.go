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
	panic(fmt.Errorf("not implemented"))
	// db, _ := pop.Connect("solomo")

	// productImages := []models.ProductImage{}
	// query := db.Where("product_id = ?", obj.ID)
	// err := query.All(&productImages)

	// return []*models_gen.ProductImage{
	// 	{
	// 		ID:    "1",
	// 		Image: "ddd",
	// 	},
	// 	{
	// 		ID:    "2",
	// 		Image: "bbb",
	// 	},
	// }, nil
	// return *productImages, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*models.Product, error) {
	// panic(fmt.Errorf("not implemented"))

	db, _ := pop.Connect("solomo")

	product := models.Product{}
	err := db.Eager("ProductImages").First(&product)

	fmt.Println(product.ID)
	fmt.Println(product.ProductImages)

	// err := db.Find(&product, 1)

	if err != nil {
		fmt.Printf("ddd %v", err)
	}

	// fmt.Println(product)

	// t := &models.Product{
	// 	No:    "NO00001",
	// 	Name:  "Apple tree",
	// 	Price: 100,
	// 	ProductImages: []models.ProductImage{
	// 		{
	// 			Image: "ddd",
	// 		},
	// 		{
	// 			Image: "ddd",
	// 		},
	// 	},
	// }

	// if err := db.Eager().Create(t); err != nil {
	// 	// return err
	// 	fmt.Println(err)
	// }

	// desp := "描述中"
	// brief := "123"

	// return &models.Product{
	// 	ID:    1,
	// 	Name:  "產品一號",
	// 	Desp:  &desp,
	// 	Brief: &brief,
	// }, nil

	return &product, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Product() generated.ProductResolver   { return &productResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
