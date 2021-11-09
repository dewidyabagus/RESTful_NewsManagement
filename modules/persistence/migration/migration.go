package migration

import (
	"gorm.io/gorm"

	// List model
	"RESTful/modules/persistence/post"
	"RESTful/modules/persistence/topic"
)

func TableMigration(db *gorm.DB) {
	db.AutoMigrate(&topic.Topic{}, &post.Post{})
}
