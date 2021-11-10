package response

import (
	"time"

	"RESTful/business/topic"
)

type PostSummary struct {
	ID          string    `json:"id"`
	TopicName   string    `json:"topic_name"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Excerpt     string    `json:"excerpt"`
	Tags        []string  `json:"tags"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TopicWithAllPosts struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Posts       []PostSummary `json:"posts"`
}

func GetAllTopic(topics *[]topic.Topic) *[]TopicDetail {
	var result = []TopicDetail{}

	for _, item := range *topics {
		result = append(result, *GetOneTopic(&item))
	}

	if result == nil {
		result = []TopicDetail{}
	}

	return &result
}

func GetTopicWithAllPosts(t *topic.TopicWithPosts) *TopicWithAllPosts {
	var result = TopicWithAllPosts{}

	result.ID = t.ID
	result.Name = t.Name
	result.Description = t.Description
	result.UpdatedAt = t.UpdatedAt

	for _, post := range *t.Posts {
		result.Posts = append(result.Posts, PostSummary{
			ID:          post.ID,
			TopicName:   post.TopicName,
			Title:       post.Title,
			Slug:        post.Slug,
			Excerpt:     post.Excerpt,
			Tags:        post.Tags,
			PublishedAt: post.PublishedAt,
			UpdatedAt:   post.UpdatedAt,
		})
	}

	if result.Posts == nil {
		result.Posts = []PostSummary{}
	}

	return &result
}
