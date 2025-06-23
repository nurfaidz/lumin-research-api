package models

type Category struct {
	GormModel
	Name string `json:"name" binding:"required" gorm:"not null"`
}
