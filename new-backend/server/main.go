package main

import (
	"fmt"
	"log"
	"strings"

	middlewares "XinAPI/app/http/middlewares"
	routesGraphql "XinAPI/routes/graphql/admin"

	_ "XinAPI/pkg/loadEnv"

	"XinAPI/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func ParseTokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Request.Header["Authorization"]; ok {
			tokenWithBearer := c.Request.Header.Get("Authorization")
			splitToken := strings.Split(tokenWithBearer, "Bearer")
			token := splitToken[1]

			c.Set("role", "yoyotv..")

			fmt.Println(token)
			c.Next()
		}
	}
}

func main() {

	port := viper.GetString("port.admin")

	// Setting up Gin
	r := gin.Default()

	r.Use(ParseTokenHandler())
	r.Use(middlewares.GinContextToContext())

	r.Use(middlewares.Cors(viper.GetStringSlice("cors.admin")...))

	routesGraphql.SetGraphQLRouter(r)

	routes.SetAPIRouter(r)

	log.Fatal(r.Run(":" + port))
}
