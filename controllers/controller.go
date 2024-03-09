package controllers

import "github.com/gin-gonic/gin"

func ShowStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "John Cena",
	})
}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		name: "Do you always come here?",
	})
}
