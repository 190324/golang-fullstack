package main

import (
	"log"

	middlewares "XinAPI/app/http/middlewares"
	routesGraphql "XinAPI/routes/graphql/api"

	_ "XinAPI/pkg/loadEnv"

	"XinAPI/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	port := viper.GetString("port.api")

	// Setting up Gin
	r := gin.Default()

	r.Use(middlewares.Cors(viper.GetStringSlice("cors.api")...))

	routesGraphql.SetGraphQLAPIRouter(r)

	routes.SetAPIRouter(r)

	log.Fatal(r.Run(":" + port))
}
