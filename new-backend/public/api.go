package main

import (
	"log"

	generated "XinAPI/app/graphql/api"
	resolvers "XinAPI/app/graphql/api/resolvers"
	middlewares "XinAPI/app/http/middlewares"

	_ "XinAPI/pkg/loadEnv"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func main() {

	port := viper.GetString("port.api")

	// Setting up Gin
	r := gin.Default()

	r.Use(middlewares.Cors(viper.GetStringSlice("cors.api")...))

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	log.Fatal(r.Run(":" + port))
}
