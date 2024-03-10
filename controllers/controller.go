package controllers

import (
	"github.com/george-prado/gin-api-rest/initializers"
	"github.com/george-prado/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	initializers.DB.Find(&students)
	c.JSON(200, students)
}

func PostStudent(c *gin.Context){
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	initializers.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func GetStudentById(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	initializers.DB.First(&student, id)
	
	if student.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found"})
		return
	}
	
	c.JSON(http.StatusOK, student)
}

func GetStudentsByAge(c *gin.Context){
	var students []models.Student
	ageStr :=  c.Param("age")
	
	if ageStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "You must provide an age"})
		return
	}
	
	initializers.DB.Where("age = ?", ageStr).Find(&students)
	
	if len(students) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Student not found"})
		return
	}
	
	c.JSON(http.StatusOK, students)
}

func UpdateStudentById(c *gin.Context){
	var student models.Student
	id := c.Params.ByName("id")
	initializers.DB.First(&student, id)
	
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	
	initializers.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, gin.H{
		"success": "student info updated"})
}

func DeleteStudentById(c *gin.Context){
	var student models.Student
	id := c.Params.ByName("id")
	initializers.DB.First(&student, id)
	
	if student.Id != 0 {
		initializers.DB.Delete(&student)
		c.JSON(http.StatusOK, gin.H{
			"Success": "Student deleted"})
		return
	}
	
	c.JSON(http.StatusNotFound, gin.H{
		"Error": "Student not found"})
}

func DeleteAll(c *gin.Context) {
	var students []models.Student
	
	initializers.DB.Find(&students)
	
	for _, student := range students {
		initializers.DB.Delete(&student)
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Cleared database"})
}