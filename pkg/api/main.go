package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getDevices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"device": "found device",
		})
	}
}

func Endpoints() {
	routes := gin.Default()

	routes.GET("/devices", getDevices())
	routes.Run()
}
