package page

import (
	"backend-intern-homework/infra/database"
	"backend-intern-homework/models"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func GetPages() (pages []models.Page, err error) {
	err = database.DB.Preload(clause.Associations).Find(&pages).Error

	return pages, err
}

func GetPage(key uuid.UUID) (page *models.Page, err error) {
	err = database.DB.Preload(clause.Associations).First(&page, key).Error

	return page, err
}

func CreatePage(page *models.Page) (err error) {
	err = database.DB.Create(&page).Error

	return err
}

func UpdatePage(page *models.Page) (err error) {
	var oldPage models.Page

	err = database.DB.First(&oldPage, page.Key).Error

	if err == nil {
		err = database.DB.Model(&oldPage).Association("Articles").Clear()

		err = database.DB.Updates(&page).Error
	}

	return err
}

func DeletePage(page *models.Page) (err error) {
	err = database.DB.First(&models.Page{}, page.Key).Error

	if err == nil {
		err = database.DB.Select(clause.Associations).Delete(&page).Error
	}

	return err
}
