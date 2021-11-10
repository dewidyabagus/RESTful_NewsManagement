package topic

import (
	"RESTful/business/post"
	"time"

	"github.com/google/uuid"
)

type Topic struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TopicWithPosts struct {
	ID          string
	Name        string
	Description string
	UpdatedAt   time.Time
	Posts       *[]post.Post
}

type TopicSpec struct {
	Name        string `validate:"required,max=75"`
	Description string `validate:"required,max=150"`
}

func (t *TopicSpec) toInsertTopic() *Topic {
	return &Topic{
		ID:          uuid.New().String(),
		Name:        t.Name,
		Description: t.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (t *TopicSpec) toUpdateTopic() *Topic {
	return &Topic{
		Name:        t.Name,
		Description: t.Description,
		UpdatedAt:   time.Now(),
	}
}

func GetTopicWithAllPosts(t *Topic, posts *[]post.Post) *TopicWithPosts {
	return &TopicWithPosts{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		UpdatedAt:   t.UpdatedAt,
		Posts:       posts,
	}
}
