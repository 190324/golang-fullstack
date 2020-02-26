package main

import (
	"log"

	middlewares "XinAPI/app/http/middlewares"
	routesGraphql "XinAPI/routes/graphql/admin"

	_ "XinAPI/pkg/loadEnv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	port := viper.GetString("port.admin")

	// Setting up Gin
	r := gin.Default()

	r.Use(middlewares.Cors(viper.GetStringSlice("cors.admin")...))

	routesGraphql.SetGraphQLRouter(r)

	log.Fatal(r.Run(":" + port))
}
