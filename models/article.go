package models

type Article struct {
	GormModel
	Title      string   `json:"title" binding:"required" gorm:"not null;type:text"`
	Slug       string   `json:"slug" gorm:"not null"`
	CategoryId uint     `json:"category_id" binding:"required" gorm:"not null"`
	UserId     uint     `json:"user_id" gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryId"`
	User       User     `gorm:"foreignKey:UserId"`
}
