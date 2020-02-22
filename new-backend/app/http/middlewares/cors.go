package middlewares

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(corsList ...string) gin.HandlerFunc {
	config := cors.DefaultConfig()

	// aInterface := data["aString"].([]interface{})
	// aString := make([]string, len(corsList))
	// for i, v := range corsList {
	// 	aString[i] = v.(string)
	// }

	if len(corsList) == 0 {
		config.AllowOrigins = []string{"http://localhost"}
	} else {
		fmt.Println(corsList[0])
		config.AllowOrigins = corsList
	}

	return cors.New(config)
}
