package routes

import (
	"github.com/george-prado/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowStudents)
	r.GET("/:name", controllers.Greetings)

	r.Run()
}
