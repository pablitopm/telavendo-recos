package server

import (
	"net/http"

	"telavendo.com.ar/telavendo-recos/app/interface/rest"

	"github.com/gin-gonic/gin"
)

func createRoutes(r *gin.Engine) {
	r.Use(InjectContainer())
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/order", rest.SaveOrderProducts)

}
