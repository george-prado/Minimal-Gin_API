package models

import (
	"time"
)

type Student struct {
	Id        uint64     `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Age        int       `gorm:"type:int;not null"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
}
