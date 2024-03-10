package migrate

import (
	"fmt"
	"github.com/george-prado/gin-api-rest/initializers"
	"log"
	
	"github.com/george-prado/gin-api-rest/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}
	
	initializers.ConnectDB(&config)
}

func MigrateDB() {
	initializers.DB.AutoMigrate(&models.Student{})
	fmt.Println("? Migration complete")
}