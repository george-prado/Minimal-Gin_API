package routes

import (
	"github.com/george-prado/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/api/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	r.GET("/students/all", controllers.GetAllStudents)
	r.GET("/students/id/:id", controllers.GetStudentById)
	r.GET("/students/age/:age", controllers.GetStudentsByAge)
	r.POST("/student/create", controllers.PostStudent)
	r.PUT("/students/update/:id", controllers.UpdateStudentById)
	r.DELETE("/students/delete/:id", controllers.DeleteStudentById)
	r.DELETE("/data/clearall", controllers.DeleteAll)
	r.Run()
}
