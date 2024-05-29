package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name          string `json:"name" form:"name" gorm:"unique"`
	Age           int    `json:"age" form:"age"`
	BirthDate     string `json:"birth_date" form:"birth_date"`
	Gender        string `json:"gender" form:"gender"`
	UserID        int    `json:"user_id" form:"user_id" gorm:"unique"`
	Participation []Participation
}
