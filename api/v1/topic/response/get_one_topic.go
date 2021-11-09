package response

import (
	"RESTful/business/topic"
	"time"
)

type TopicDetail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func GetOneTopic(topic *topic.Topic) *TopicDetail {
	return &TopicDetail{
		ID:          topic.ID,
		Name:        topic.Name,
		Description: topic.Description,
		UpdatedAt:   topic.UpdatedAt,
	}
}
