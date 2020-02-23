// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package resolvers

import (
	generated "XinAPI/app/graphql/api"
	models_gen "XinAPI/app/graphql/api/models"
	"XinAPI/app/http/middlewares"
	"XinAPI/app/mail"
	"XinAPI/pkg/jwt"
	"fmt"
	"time"

	// . "XinAPI/pkg/log"
	"context"
)

func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (*string, error) {
	// panic(fmt.Errorf("not implemented"))

	var str = "success"

	mail.ForgotPassword(email)

	return &str, nil
}

func (r *mutationResolver) Login(ctx context.Context, input models_gen.InputLogin) (*models_gen.PayloadAuth, error) {

	if input.Email == "test@liontravel.com" && input.Password == "0000" {

		validTime := time.Hour * time.Duration(1)
		jwtdata := jwt.JWTStandard{
			Id:        "1",
			Name:      "hello world",
			ValidTime: validTime,
			Issuer:    "xinmedia.com",
			Subject:   "login",
		}

		token, _ := jwtdata.Create()

		return &models_gen.PayloadAuth{
			AccessToken:  token,
			RefreshToken: token,
			ExpiresIn:    int(time.Now().Add(validTime).Unix()),
			TokenType:    "login",
		}, nil
	}

	return nil, nil
}

func (r *queryResolver) Me(ctx context.Context) (*models_gen.Member, error) {
	// panic(fmt.Errorf("not implemented"))

	gc, _ := middlewares.GinContextFromContext(ctx)

	fmt.Println(gc.Get("role"))

	return &models_gen.Member{
		ID:   "123",
		Name: "hello world",
	}, nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
