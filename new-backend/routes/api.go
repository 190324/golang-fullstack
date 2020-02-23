package routes

import (
	"XinAPI/app/http/controllers"

	"github.com/gin-gonic/gin"
)

func SetAPIRouter(r *gin.Engine) {
	router := r.Group("api")

	router.GET("/health", new(controllers.Home).Index)

	return
}
