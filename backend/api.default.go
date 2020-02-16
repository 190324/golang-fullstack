package main

import (
	"log"
	"net/http"
	"os"

	generated "example.com/gqlgen/app/graphQL"
	resolvers "example.com/gqlgen/app/graphQL/resolvers"
	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
