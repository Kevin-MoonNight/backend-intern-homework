package migrations

import (
	"backend-intern-homework/infra/database"
	"backend-intern-homework/infra/logger"
	"backend-intern-homework/models"
)

// Migrate Add list of model add for migrations
func Migrate() {
	var migrationModels = []interface{}{&models.Article{}, &models.Page{}, &models.List{}}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		logger.Errorf("Migration error: %s", err)
	}
}
