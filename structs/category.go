package structs

type CategoryResponse struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required" gorm:"unique;not null"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"required" gorm:"unique:not null"`
}
