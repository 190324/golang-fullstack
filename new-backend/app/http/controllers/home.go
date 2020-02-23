package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Home struct {
	Aac string
}

func (ctl *Home) Index(c *gin.Context) {
	c.JSON(http.StatusOK, "health")
}
