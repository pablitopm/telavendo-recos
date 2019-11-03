package server

import (
	"telavendo.com.ar/telavendo-recos/app/registry"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InjectContainer() gin.HandlerFunc {
	log.Debug("Starting container")
	ctn, err := registry.NewContainer()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}
	return func(c *gin.Context) {
		c.Set("ctn", ctn)
		c.Next()
	}
}
