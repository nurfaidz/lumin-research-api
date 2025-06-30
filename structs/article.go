package structs

type ArticleResponse struct {
	Id         uint   `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	CategoryId uint   `json:"category_id"`
	UserId     uint   `json:"user_id"`
}

type ArticleCreateRequest struct {
	Title      string `json:"title" binding:"required" gorm:"not null;type:text"`
	Slug       string `json:"slug" gorm:"not null"`
	CategoryId uint   `json:"category_id" binding:"required" gorm:"not null"`
	UserId     uint   `json:"user_id" gorm:"not null"`
}

type ArticleUpdateRequest struct {
	Title      string `json:"title" binding:"required" gorm:"not null;type:text"`
	Slug       string `json:"slug" gorm:"not null"`
	CategoryId uint   `json:"category_id" binding:"required" gorm:"not null"`
	UserId     uint   `json:"user_id" gorm:"not null"`
}
