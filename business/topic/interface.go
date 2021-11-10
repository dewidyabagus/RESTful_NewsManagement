package topic

import "time"

type Service interface {
	// Insert new topic
	InsertTopic(topic *TopicSpec) error

	// Find topic by id (uuid)
	FindTopicById(id *string) (*Topic, error)

	// Find topic and all post
	FindTopicByNameWithAllPosts(name *string) (*TopicWithPosts, error)

	// Find all topic
	FindAllTopic() (*[]Topic, error)

	// Update topic information
	UpdateTopic(id *string, topic *TopicSpec) error

	// Delete topic with soft delete
	DeleteTopic(id *string) error
}

type Repository interface {
	// Insert new topic
	InsertTopic(topic *Topic) error

	// Find topic by name, name topic unique
	FindTopicByName(name *string) (*Topic, error)

	// Find topic by id (uuid)
	FindTopicById(id *string) (*Topic, error)

	// Find all topic
	FindAllTopic() (*[]Topic, error)

	// Update topic information
	UpdateTopic(id *string, topic *Topic) error

	// Deleted topic with soft delete
	DeleteTopic(id *string, deleter time.Time) error
}
