package models

type Article struct {
	ID      uint   `json:"-" gorm:"primarykey;"`
	Title   string `json:"title" gorm:"not null;"`
	Content string `json:"content" gorm:"not null;"`
}

type CreateArticle struct {
	Title   string `json:"title" validate:"required,max=256"`
	Content string `json:"content" validate:"required"`
}

type UpdateArticle struct {
	Title   string `json:"title" validate:"required,max=256"`
	Content string `json:"content" validate:"required"`
}
