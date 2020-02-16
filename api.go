package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	generated "example.com/gqlgen/app/graphQL"
	resolvers "example.com/gqlgen/app/graphQL/resolvers"
	middlewares "example.com/gqlgen/app/http/middlewares"
	models_gen "example.com/gqlgen/graphql"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	c := generated.Config{Resolvers: &resolvers.Resolver{}}
	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models_gen.Role) (interface{}, error) {
		// if !getCurrentUser(ctx).HasRole(role) {
		// 	// block calling the next resolver
		// 	return nil, fmt.Errorf("Access denied")
		// }
		return nil, fmt.Errorf("Access denied")

		// or let it pass through
		return next(ctx)
	}

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

const defaultPort = "8080"

func main() {

	port := defaultPort

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		// os.Exit(1)
	}

	port = os.Getenv("PORT")

	// Setting up Gin
	r := gin.Default()
	r.Use(middlewares.CorsMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	log.Fatal(r.Run(":" + port))
}
