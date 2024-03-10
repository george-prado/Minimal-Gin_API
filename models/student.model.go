package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Students []Student
