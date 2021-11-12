package main

import (
	"fmt"

	"github.com/go-redis/redis"
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

	// Topic Caching
	topicServiceCache "RESTful/business/cache/topic"
	topicRepoCache "RESTful/modules/cache/topic"

	// Post
	postController "RESTful/api/v1/post"
	postService "RESTful/business/post"
	postRepository "RESTful/modules/persistence/post"

	// Post Caching
	postServiceCache "RESTful/business/cache/post"
	postRepoCache "RESTful/modules/cache/post"
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

func newRedisConnection(config *config.AppConfig) *redis.Client {
	rdClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       0,
	})

	if _, err := rdClient.Ping().Result(); err != nil {
		panic(err)
	}

	return rdClient
}

func main() {
	// Get configuration file
	config := config.GetAppConfig()

	// Create new session database
	dbConnection := newDatabaseConnection(config)

	// Create ne session redis
	rdClient := newRedisConnection(config)

	// Initiate topic repository redis
	topicRCache := topicRepoCache.NewRepository(rdClient)

	// Initiate topic repository
	topicRepo := topicRepository.NewRepository(dbConnection)

	// Initiate topic service cache
	topicSCache := topicServiceCache.NewService(topicRCache)

	// Initiate post repository
	postRepo := postRepository.NewRepository(dbConnection)

	// Initiate post cache repository
	postRCache := postRepoCache.NewRepository(rdClient)

	// Initiate post cache service
	postSCache := postServiceCache.NewService(postRCache)

	// Initiate post service
	postSvc := postService.NewService(postRepo, postSCache)

	// Initiate post controller
	postHandler := postController.NewController(postSvc)

	// Initiate topic service
	topicSvc := topicService.NewService(topicRepo, postSvc, topicSCache)

	// Initiate topic controller
	topicHandler := topicController.NewController(topicSvc)

	// Initiate echo web framework
	e := echo.New()

	// Initiate routes
	api.RegisterRouters(e, topicHandler, postHandler)

	// start echo
	e.Start(fmt.Sprintf("%s:%d", config.AppHost, config.AppPort))
}
