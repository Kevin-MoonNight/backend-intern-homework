package list

import (
	"backend-intern-homework/infra/database"
	"backend-intern-homework/models"

	"github.com/google/uuid"
)

func GetLists() (lists []models.List, err error) {
	err = database.DB.Find(&lists).Error

	return lists, err
}

func GetList(key uuid.UUID) (list *models.List, err error) {
	err = database.DB.First(&list, key).Error

	return list, err
}

func CreateList(list *models.List) (err error) {
	err = database.DB.Create(&list).Error

	return err
}

func UpdateList(list *models.List) (err error) {
	err = database.DB.First(&models.List{}, list.Key).Error

	if err == nil {
		err = database.DB.Updates(&list).Error
	}

	return err
}

func DeleteList(list *models.List) (err error) {
	err = database.DB.First(&models.List{}, list.Key).Error

	if err == nil {
		err = database.DB.Delete(&list).Error
	}

	return err
}
