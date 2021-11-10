package post

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID              string
	TopicId         string
	TopicName       string
	TopicDesciption string
	Title           string
	Slug            string
	Excerpt         string
	Body            string
	Tags            []string
	PublishedAt     time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type PostSpec struct {
	TopicId string   `validate:"required"`
	Title   string   `validate:"required"`
	Slug    string   `validate:"required"`
	Excerpt string   `validate:"required"`
	Body    string   `validate:"required"`
	Tags    []string `validate:"required"`
}

func (p *PostSpec) toInsertPost() *Post {
	return &Post{
		ID:        uuid.New().String(),
		TopicId:   p.TopicId,
		Title:     p.Title,
		Slug:      p.Slug,
		Excerpt:   p.Excerpt,
		Body:      p.Body,
		Tags:      p.Tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
