package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type List struct {
	Key         uuid.UUID  `json:"key" gorm:"primaryKey;type:uuid;"`
	Name        string     `json:"name"`
	NextPageKey *uuid.UUID `json:"next_page_key"`
}

type CreateList struct {
	Name        string `json:"name" validate:"required"`
	NextPageKey string `json:"next_page_key" validate:"omitempty,uuid4"`
}

type UpdateList struct {
	Name        string `json:"name" validate:"required"`
	NextPageKey string `json:"next_page_key" validate:"omitempty,uuid4"`
}

func (l *List) BeforeCreate(tx *gorm.DB) (err error) {
	key, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	l.Key = key

	return
}
