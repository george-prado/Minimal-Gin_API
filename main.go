package main

import (
	"github.com/george-prado/gin-api-rest/initializers"
	"github.com/george-prado/gin-api-rest/migrate"
	"github.com/george-prado/gin-api-rest/routes"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	
	initializers.ConnectDB(&config)
	
	
	server = gin.Default()
}

func main() {
	migrate.MigrateDB()
	
	routes.HandleRequests()
}