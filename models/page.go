package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Page struct {
	Key         uuid.UUID `json:"key" gorm:"primaryKey;type:uuid;"`
	NextPageKey uuid.UUID `json:"next_page_key"`
	Articles    []Article `json:"articles" gorm:"many2many:page_articles;"`
}

type CreatePage struct {
	NextPageKey string `json:"next_page_key" validate:"omitempty,uuid4"`
	ArticleIds  []uint `json:"article_ids" validate:"required"`
}

type UpdatePage struct {
	NextPageKey string `json:"next_page_key" validate:"omitempty,uuid4"`
	ArticleIds  []uint `json:"article_ids" validate:"required"`
}

func (p *Page) BeforeCreate(tx *gorm.DB) (err error) {
	key, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	p.Key = key

	return
}
