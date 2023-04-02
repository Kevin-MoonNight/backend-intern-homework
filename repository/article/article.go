package article

import (
	"backend-intern-homework/infra/database"
	"backend-intern-homework/models"
)

func GetArticles(conds ...interface{}) (articles []models.Article, err error) {
	if len(conds) == 0 {
		err = database.DB.Find(&articles).Order("id DESC").Error
	} else {
		err = database.DB.Find(&articles, conds[0]).Order("id DESC").Error
	}

	return articles, err
}

func FindArticle(id uint) (article *models.Article, err error) {
	err = database.DB.First(&article, id).Error

	return article, err
}

func CreateArticle(article *models.Article) (err error) {
	err = database.DB.Create(&article).Error

	return err
}

func UpdateArticle(article *models.Article) (err error) {
	err = database.DB.First(&models.Article{}, article.ID).Error

	if err == nil {
		err = database.DB.Updates(&article).Error
	}

	return err
}

func DeleteArticle(article *models.Article) (err error) {
	err = database.DB.First(&models.Article{}, article.ID).Error

	if err == nil {
		err = database.DB.Delete(&article).Error
	}

	return err
}
