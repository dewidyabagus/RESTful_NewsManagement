package response

import (
	"RESTful/business/post"
	"time"
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

func GetOnePostSummary(p *post.Post) *PostSummary {
	return &PostSummary{
		ID:          p.ID,
		TopicName:   p.TopicName,
		Title:       p.Title,
		Slug:        p.Slug,
		Excerpt:     p.Excerpt,
		Tags:        p.Tags,
		PublishedAt: p.PublishedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
