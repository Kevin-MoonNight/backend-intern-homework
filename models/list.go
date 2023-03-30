package models

type List struct {
	Key         string `gorm:"primaryKey;"`
	NextPageKey string `json:"next_page_key"`
}
