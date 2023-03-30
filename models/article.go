package models

type Article struct {
	ID      uint   `json:"-" gorm:"primarykey;"`
	Title   string `json:"title" gorm:"not null;"`
	Content string `json:"content" gorm:"not null;"`
}
