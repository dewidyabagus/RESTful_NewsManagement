package request

import (
	"RESTful/business/topic"
)

type Topic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *Topic) ToBusinessTopicSpec() *topic.TopicSpec {
	return &topic.TopicSpec{
		Name:        t.Name,
		Description: t.Description,
	}
}
