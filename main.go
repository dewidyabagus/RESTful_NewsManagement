package main

import (
	// Module configuration file
	"RESTful/config"

	// Migration
	"RESTful/modules/migration"

	// API
	"RESTful/api"

	"fmt"

	echo "github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabaseConnection(config *config.AppConfig) *gorm.DB {
	strConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.PgHost, config.PgUsername, config.PgPassword, config.PgDbname, config.PgPort,
	)

	db, err := gorm.Open(postgres.Open(strConnection), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func main() {
	// Get configuration file
	config := config.GetAppConfig()

	// Create new session database
	_ = newDatabaseConnection(config)

	// Initiate echo web framework
	e := echo.New()

	// Initiate routes
	api.RegisterRouters(e)

	// start echo
	e.Start(":8000")
}
