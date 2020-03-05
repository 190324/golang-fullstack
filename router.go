package router

import (
	resolvers "XinAPI/app/graphql/admin/resolvers"
	generated "XinAPI/build/gqlgen/admin"
	models_gen "XinAPI/build/gqlgen/admin/models"
	"XinAPI/pkg/l"
	"github.com/99designs/gqlgen/graphql"

	"context"

	//"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	Config := generated.Config{Resolvers: &resolvers.Resolver{}}
	Config.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models_gen.Role) (interface{}, error) {
		l.Log(`AdminQuerylHandler`, role)

		// if userRole != role {
		// 	return nil, fmt.Errorf("Access denied")
		// }

		return next(ctx)
	}


	h := handler.NewDefaultServer(generated.NewExecutableSchema(Config))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func SetGraphQLRouter(r *gin.Engine) {

	router := r.Group(`graphql`)

	router.POST("/", graphqlHandler())
	router.GET("/playground", playgroundHandler())

}
