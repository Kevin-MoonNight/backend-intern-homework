package main

import (
	"backend-intern-homework/config"
	"backend-intern-homework/infra/database"
	"backend-intern-homework/infra/logger"
	"backend-intern-homework/migrations"
	"backend-intern-homework/routers"
	"time"

	"github.com/spf13/viper"
)

// @title Backend Intern Homework API
// @version 1.0
// @description This is the Backend Intern Homework api documentation
func main() {

	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Taipei")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	dbDSN := config.DBConfiguration()

	if err := database.DbConnection(dbDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	//later separate migration
	migrations.Migrate()

	router := routers.SetupRoute()
	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
