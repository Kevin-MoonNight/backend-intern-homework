package article

import (
	"backend-intern-homework/infra/database"
	"backend-intern-homework/models"
)

func GetArticles() (Articles []models.Article, err error) {
	err = database.DB.Find(&Articles).Order("id DESC").Error

	return Articles, err
}

func FindArticle(id uint) (Article *models.Article, err error) {
	err = database.DB.First(&Article, id).Error

	return Article, err
}

func CreateArticle(Article *models.Article) (err error) {
	err = database.DB.Create(&Article).Error

	return err
}

func UpdateArticle(Article *models.Article) (err error) {
	err = database.DB.First(&models.Article{}, Article.ID).Error

	if err == nil {
		err = database.DB.Updates(&Article).Error
	}

	return err
}

func DeleteArticle(Article *models.Article) (err error) {
	err = database.DB.First(&models.Article{}, Article.ID).Error

	if err == nil {
		err = database.DB.Delete(&Article).Error
	}

	return err
}
