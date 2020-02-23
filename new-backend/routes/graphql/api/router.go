package routesGraphqlAdmin

import (
	generated "XinAPI/app/graphql/api"
	resolvers "XinAPI/app/graphql/api/resolvers"

	// models_gen "XinAPI/app/graphql/api/models"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := generated.Config{Resolvers: &resolvers.Resolver{}}

	h := handler.GraphQL(generated.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func SetGraphQLAPIRouter(r *gin.Engine) {

	router := r.Group(`graphql`)

	router.POST("/", graphqlHandler())
	router.GET("/playground", playgroundHandler())

}
