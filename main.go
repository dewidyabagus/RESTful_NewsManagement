package main

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	// Module configuration file
	"RESTful/config"

	// Migration
	"RESTful/modules/persistence/migration"

	// API
	"RESTful/api"

	// Topic
	topicController "RESTful/api/v1/topic"
	topicService "RESTful/business/topic"
	topicRepository "RESTful/modules/persistence/topic"
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
	dbConnection := newDatabaseConnection(config)

	// Initiate topic repository
	topicRepo := topicRepository.NewRepository(dbConnection)

	// Initiate topic service
	topicSvc := topicService.NewService(topicRepo)

	// Initiate topic controller
	topicHandler := topicController.NewController(topicSvc)

	// Initiate echo web framework
	e := echo.New()

	// Initiate routes
	api.RegisterRouters(e, topicHandler)

	// start echo
	e.Start(":8000")
}
