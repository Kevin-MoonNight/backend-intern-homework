package models

type Page struct {
	Key         string     `gorm:"primaryKey;"`
	NextPageKey string     `json:"next_page_key"`
	Articles    []*Article `json:"articles" gorm:"many2many:page_articles;"`
}
