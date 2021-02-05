package main

import (
	"fmt"
	"net/http"

	"github.com/easyms/archv1/pc-service/internal"
	"github.com/easyms/archv1/pc-service/internal/db"
	"github.com/easyms/archv1/pc-service/pkg/routes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	config := internal.LoadConfig(".env")
	database := db.NewConnection(config)
	logger := log.New()

	defer database.Close()

	router := gin.Default()

	api := router.Group("/api")
	{
		routes.ApplyRoutes(api, database, logger)
	}
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	port := config.GetConfigProperty("SERVICE_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
