package models

// Article article model
type Article struct {
	BaseModel
	ID      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  int    `json:"status" gorm:"default:1"`
}

// TableName table name
func (art *Article) TableName() string {
	return "article"
}
