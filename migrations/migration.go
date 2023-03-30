package migrations

import (
	"backend-intern-homework/infra/database"
	"backend-intern-homework/infra/logger"
)

// Migrate Add list of model add for migrations
func Migrate() {
	var migrationModels = []interface{}{}
	err := database.DB.AutoMigrate(migrationModels...)
	if err != nil {
		logger.Errorf("Migration error: %s", err)
	}
}
