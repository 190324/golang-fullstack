package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	generated "example.com/gqlgen/app/graphQL"
	resolvers "example.com/gqlgen/app/graphQL/resolvers"
)

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))

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
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	// r.Run()
	log.Fatal(r.Run(":" + port))
}
